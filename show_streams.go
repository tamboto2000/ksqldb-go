package ksqldbx

import "context"

type Stream SourceDesc

func (s Stream) IsExtended() bool {
	return s.Extended
}

// ShowStreams get all existing streams.
// See https://docs.ksqldb.io/en/latest/developer-guide/ksqldb-rest-api/ksql-endpoint
// and https://docs.ksqldb.io/en/latest/developer-guide/ksqldb-reference/show-streams/
// for more info
func (ksql *KsqlDB) ShowStreams(ctx context.Context) ([]Stream, error) {
	var streams []Stream
	q := "SHOW STREAMS;"
	res, err := ksql.Exec(ctx, StmntSQL{KSQL: q})
	if err != nil {
		return streams, nil
	}

	return res[0].Streams, nil
}
