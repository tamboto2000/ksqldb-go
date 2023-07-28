package ksqldbx

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
)

// Pull run PULL query, which will return all result at once.
// A PULL query is a query that didn't EMIT CHANGES,
// running a query that use EMIT CHANGES will cause this method to
// block indefinitely
func (ksql *KsqlDB) Pull(ctx context.Context, q QuerySQL) (Header, []Row, error) {
	// TODO
	// check if sql parse option is enabled, if yes
	// then parse sql before exec

	var header Header
	var rows []Row

	res, err := ksql.queryStreamRequest(ctx, q)
	if err != nil {
		return header, rows, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return header, rows, newErrFromReader(res.Body)
	}

	reader := bufio.NewReader(res.Body)
	headerFound := false
	for {
		body, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF && err != context.Canceled {
				return header, rows, err
			}

			return header, rows, nil
		}

		if !headerFound {
			if err := json.Unmarshal(body, &header); err != nil {
				return header, rows, err
			}

			headerFound = true
			continue
		}

		var row Row
		if err := json.Unmarshal(body, &row); err != nil {
			return header, rows, err
		}

		rows = append(rows, row)
	}
}
