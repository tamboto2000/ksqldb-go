package ksqldbx

import "context"

type StreamExtended SourceDesc

type Stream struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Topic       string `json:"topic"`
	KeyFormat   string `json:"keyFormat"`
	ValueFormat string `json:"valueFormat"`
	IsWindowed  bool   `json:"isWindowed"`
}

// ShowStreams fetch all existing streams
func (ksql *KsqlDB) ShowStreams(ctx context.Context) ([]Stream, error) {
	var streams []Stream
	q := "SHOW STREAMS;"
	res, err := ksql.Exec(ctx, StmntSQL{KSQL: q})
	if err != nil {
		return streams, nil
	}

	return res[0].Streams, nil
}
