package ksqldbx

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamboto2000/ksqldbx/net"
)

func TestIntegrationKsqlDB_ShowStreams(t *testing.T) {
	createStreamStmnt := `CREATE STREAM IF NOT EXISTS show_streams_test (
		d1 BOOLEAN
	) WITH (
		kafka_topic='test_stream', 
		value_format='protobuf',
		partitions=1
	);`

	ksql, err := NewKsqlDB(net.Options{
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	})

	require.Nil(t, err)

	// create a stream
	_, err = ksql.Exec(context.Background(), StmntSQL{
		KSQL: createStreamStmnt,
	})

	require.Nil(t, err)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name        string
		args        args
		wantStreams []Stream
		wantErr     bool
	}{
		{
			name: "success get list of streams",
			args: args{
				ctx: context.Background(),
			},
			wantStreams: []Stream{
				{
					Type:        "STREAM",
					Name:        "SHOW_STREAMS_TEST",
					Topic:       "test_stream",
					KeyFormat:   "KAFKA",
					ValueFormat: "PROTOBUF",
					IsWindowed:  false,
				},
				{
					Type:        "STREAM",
					Name:        "KSQL_PROCESSING_LOG",
					Topic:       "default_ksql_processing_log",
					KeyFormat:   "KAFKA",
					ValueFormat: "JSON",
					IsWindowed:  false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		streams, err := ksql.ShowStreams(tt.args.ctx)
		assert.Equal(t, tt.wantStreams, streams)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}

	// delete stream
	_, err = ksql.Exec(context.Background(), StmntSQL{
		KSQL: fmt.Sprintf("DROP STREAM %s;", "show_streams_test"),
	})

	require.Nil(t, err)
}
