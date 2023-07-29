package ksqldbx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamboto2000/ksqldbx/net"
)

func TestIntegrationKsqlDB_CloseQuery(t *testing.T) {
	createStreamStmnt := `CREATE STREAM close_query_test (
		d1 VARCHAR
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

	// drop stream
	defer func() {
		_, err := ksql.Exec(context.Background(), StmntSQL{KSQL: "DROP STREAM close_query_test;"})
		require.Nil(t, err)
	}()

	// create stream
	_, err = ksql.Exec(context.Background(), StmntSQL{KSQL: createStreamStmnt})
	require.Nil(t, err)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		args       args
		getqueryId func(ctx context.Context, ksql *KsqlDB) string
		wantErr    bool
	}{
		{
			name: "success closing a query",
			args: args{
				ctx: context.Background(),
			},
			getqueryId: func(ctx context.Context, ksql *KsqlDB) string {
				headChan := make(chan Header)
				rowChan := make(chan Row)
				go func() {
					err := ksql.Push(ctx, QuerySQL{SQL: "SELECT * FROM close_query_test EMIT CHANGES;"}, headChan, rowChan)
					require.Nil(t, err)
				}()

				head := <-headChan
				return head.QueryID
			},
			wantErr: false,
		},
		{
			name: "fail closing non existing query",
			args: args{
				ctx: context.Background(),
			},
			getqueryId: func(ctx context.Context, ksql *KsqlDB) string {
				return "abc"
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		qId := tt.getqueryId(tt.args.ctx, ksql)
		err = ksql.CloseQuery(tt.args.ctx, qId)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
