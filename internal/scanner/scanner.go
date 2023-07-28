package scanner

import (
	"errors"
	"fmt"
	"reflect"
)

// Scan scan val to dest condition that dest is a pointer and
// val data type is same or can convert to dest underlying value data type
func Scan(dest, val any) error {
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

	if val == nil {
		dv.Set(reflect.Zero(dv.Type()))
		return nil
	}

	vv := reflect.ValueOf(val)
	for vv.Kind() == reflect.Interface || vv.Kind() == reflect.Pointer {
		vv = vv.Elem()
	}

	if !vv.CanConvert(dv.Type()) {
		return fmt.Errorf("cannot scan value (type of %s) into destination (type of %s)", vv.Type().String(), dv.Type().String())
	}

	vv = vv.Convert(dv.Type())
	dv.Set(vv)

	return nil
}
