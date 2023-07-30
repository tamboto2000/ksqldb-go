package ksqldbx

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"io"
)

var ErrNoRows = errors.New("no rows in result set")

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

			if len(rows) == 0 {
				return header, rows, ErrNoRows
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

// PullRow pull exactly one Row of data.
// It does not care even if you did not specify a LIMIT 1,
// it will only return the first Row of the result set
func (ksql *KsqlDB) PullRow(ctx context.Context, q QuerySQL) (Header, Row, error) {
	var header Header
	var row Row

	header, rows, err := ksql.Pull(ctx, q)
	if err != nil {
		return header, row, err
	}

	return header, rows[0], nil
}
