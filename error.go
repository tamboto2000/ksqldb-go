package ksqldbx

import (
	"encoding/json"
	"fmt"
	"io"
)

type Error struct {
	Type          string `json:"@type"`
	ErrorCode     int    `json:"error_code"`
	Message       string `json:"message"`
	StatementText string `json:"statementText"`
	Entities      []any  `json:"entities"`
}

func (err Error) Error() string {
	return fmt.Sprintf("error code %d: %s", err.ErrorCode, err.Message)
}

func newErrFromReader(reader io.Reader) error {
	var err Error
	if err := json.NewDecoder(reader).Decode(&err); err != nil {
		return err
	}

	return err
}
