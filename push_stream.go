package ksqldbx

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

type PushStream struct {
	ksql     *KsqlDB
	ctx      context.Context
	q        QuerySQL
	header   Header
	rowChan  chan Row
	nextChan chan bool
	row      Row
	err      error
}

func errScanFieldMismatchType(colName, fieldName, colType, fieldType string) error {
	return fmt.Errorf("error on scanning column %s to destination %s: column is type of %s while destination is type of %s", colName, fieldName, colType, fieldType)
}

func (ksql *KsqlDB) NewPushStream(ctx context.Context, q QuerySQL) *PushStream {
	pushStream := &PushStream{
		ksql:     ksql,
		ctx:      ctx,
		q:        q,
		rowChan:  make(chan Row),
		nextChan: make(chan bool),
	}

	return pushStream
}

func (p *PushStream) StartStream() {
	go p.stream()
}

func (p *PushStream) stream() {
	headChan := make(chan Header)
	go func() {
		p.header = <-headChan
		for row := range p.rowChan {
			p.row = row
			p.nextChan <- true
		}
	}()

	err := p.ksql.Push(p.ctx, p.q, headChan, p.rowChan)
	if err != nil {
		p.stop(err)
	}

	p.stop(nil)
}

func (p *PushStream) stop(err error) {
	p.err = err
	p.nextChan <- false
	close(p.nextChan)
}

func (p *PushStream) Header() Header {
	return p.header
}

func (p *PushStream) Err() error {
	return p.err
}

func (p *PushStream) Next() bool {
	return <-p.nextChan
}

func (p *PushStream) Scan(dest ...any) error {
	return p.row.Scan(dest...)
}

func (p *PushStream) ScanStruct(dest any) error {
	if dest == nil {
		return errors.New("destination is nil")
	}

	dv := reflect.ValueOf(dest)
	isPointer := false
	for dv.Kind() == reflect.Interface || dv.Kind() == reflect.Pointer {
		if dv.Kind() == reflect.Pointer {
			isPointer = true
		}

		dv = dv.Elem()
	}

	if !isPointer {
		return errors.New("destination is not pointer")
	}

	if dv.Kind() != reflect.Struct {
		return errors.New("destination is not struct")
	}

	for colIdx, colName := range p.header.ColumnNames {
		rowVal := p.row[colIdx]
		colType := p.header.ColumnTypes[colIdx]

		for i := 0; i < dv.NumField(); i++ {
			f := dv.Field(i)
			if f.IsValid() {
				ft := dv.Type().Field(i)
				fName := ft.Tag.Get("ksql")
				if strings.ToUpper(fName) != colName {
					continue
				}

				for f.Kind() == reflect.Interface || f.Kind() == reflect.Pointer {
					f = f.Elem()
				}

				switch f.Kind() {
				case reflect.Bool:
					if colType != "BOOLEAN" {
						return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
					}

					f.SetBool(rowVal.(bool))

				case reflect.String:
					if colType != "STRING" {
						return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
					}

					f.SetString(rowVal.(string))

				case reflect.Slice:
					if colType == "BYTES" {
						if !f.CanConvert(reflect.TypeOf([]byte{})) {
							return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
						}

						b, err := base64.StdEncoding.DecodeString(rowVal.(string))
						if err != nil {
							return err
						}

						f.SetBytes(b)
					} else if subMatch := regexp.MustCompile(`^ARRAY<(.+)>$`).FindStringSubmatch(colType); len(subMatch) != 0 {
						arrType := subMatch[1]
						fElemKind := f.Type().Elem().Kind()
						rowValSlice := rowVal.([]interface{})
						switch arrType {
						case "BOOLEAN":
							if fElemKind != reflect.Bool {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							}

							for _, r := range rowValSlice {
								f.Set(reflect.Append(f, reflect.ValueOf(r.(bool))))
							}

						case "STRING":
							if fElemKind != reflect.String {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							}

							for _, r := range rowValSlice {
								f.Set(reflect.Append(f, reflect.ValueOf(r.(string))))
							}

						case "BYTES":
							// field must be of type [][]byte or [][]uint8
							if fElemKind != reflect.Slice {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							} else {
								// since this is 2D slice, we need to call Elem() twice
								if f.Type().Elem().Elem().Kind() != reflect.Uint8 {
									return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
								}
							}

							for _, r := range rowValSlice {
								b, err := base64.StdEncoding.DecodeString(r.(string))
								if err != nil {
									return err
								}

								f.Set(reflect.Append(f, reflect.ValueOf(b)))
							}
						case "INTEGER", "BIGINTEGER", "DOUBLE":
							// any numeric value that we get from ksqlDB will be in type of float64,
							// but since float64 can be converted into any integer type,
							// we just need to check if f's type is compatible with float64
							floatType := reflect.ValueOf(float64(0))
							if !floatType.CanConvert(f.Type().Elem()) {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							}

							for _, r := range rowValSlice {
								f.Set(reflect.Append(f, reflect.ValueOf(r).Convert(f.Type())))
							}
						case "TIME":
							if fElemKind == reflect.String {
								for _, r := range rowValSlice {
									f.Set(reflect.Append(f, reflect.ValueOf(r.(string))))
								}

								continue
							}

							refT := time.Time{}
							zero := reflect.Zero(f.Type().Elem())
							if !zero.CanConvert(reflect.TypeOf(refT)) {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							}

							for _, r := range rowValSlice {
								t, err := time.Parse("15:04:05", r.(string))
								if err != nil {
									return err
								}

								f.Set(reflect.Append(f, reflect.ValueOf(t)))
							}
						case "DATE":
							if fElemKind == reflect.String {
								for _, r := range rowValSlice {
									f.Set(reflect.Append(f, reflect.ValueOf(r.(string))))
								}

								continue
							}

							refT := time.Time{}
							zero := reflect.Zero(f.Type().Elem())
							if !zero.CanConvert(reflect.TypeOf(refT)) {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							}

							for _, r := range rowValSlice {
								t, err := time.Parse("2006-01-02", r.(string))
								if err != nil {
									return err
								}

								f.Set(reflect.Append(f, reflect.ValueOf(t)))
							}

						case "TIMESTAMP":
							if fElemKind == reflect.String {
								for _, r := range rowValSlice {
									f.Set(reflect.Append(f, reflect.ValueOf(r.(string))))
								}

								continue
							}

							refT := time.Time{}
							zero := reflect.Zero(f.Type().Elem())
							if !zero.CanConvert(reflect.TypeOf(refT)) {
								return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
							}

							for _, r := range rowValSlice {
								t, err := time.Parse("2006-01-02T15:04:05.000", r.(string))
								if err != nil {
									return err
								}

								f.Set(reflect.Append(f, reflect.ValueOf(t)))
							}
						default:
							// try to scan for DECIMAL type
							rgx := regexp.MustCompile(`^DECIMAL\(([0-9]+), ([0-9]+)\)$`)
							subMatch := rgx.FindStringSubmatch(arrType)
							if len(subMatch) != 0 {
								floatType := reflect.ValueOf(float64(0))
								if !floatType.CanConvert(f.Type().Elem()) {
									return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
								}

								for _, r := range rowValSlice {
									f.Set(reflect.Append(f, reflect.ValueOf(r).Convert(f.Type())))
								}

								continue
							}
						}

					} else {
						return errScanFieldMismatchType(colName, fName, colType, f.Type().String())
					}
				}
			}
		}
	}

	return nil
}

func (p *PushStream) scanField(val any, dest reflect.Value) {

}

func (p *PushStream) Row() Row {
	return Row{}
}

func (p *PushStream) Stop() {
	err := p.ksql.CloseQuery(p.ctx, p.header.QueryID)
	p.stop(err)
}
