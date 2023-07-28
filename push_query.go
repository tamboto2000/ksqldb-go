package ksqldbx

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// Push run PUSH query that will continuously stream data through channel.
// To stop the query you have two ways: cancel the context, or call KsqlDB.CloseQuery(),
// both methods give the same effect, it will stop the query and close the channels.
// A query is a PUSH query if you use EMIT CHANGES in the query, otherwise you might
// get unexpected behavior or error
func (ksql *KsqlDB) Push(ctx context.Context, q QuerySQL, headChan chan<- Header, rowChan chan<- Row) error {
	defer close(headChan)
	defer close(rowChan)

	// TODO
	// check if sql parse option is enabled, if yes
	// then parse sql before exec

	res, err := ksql.queryStreamRequest(ctx, q)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return newErrFromReader(res.Body)
	}

	reader := bufio.NewReader(res.Body)

	var header *Header

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			body, err := reader.ReadBytes('\n')
			if err != nil {
				if err != io.EOF && err != context.Canceled {
					return err
				}

				return nil
			}

			if len(body) == 0 {
				continue
			}

			if header == nil {
				header = new(Header)
				if err := json.Unmarshal(body, header); err != nil {
					return err
				}

				headChan <- *header
				continue
			}

			var row Row
			if err := json.Unmarshal(body, &row); err != nil {
				return err
			}

			rowChan <- row
		}
	}
}
