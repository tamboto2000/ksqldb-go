package ksqldbx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamboto2000/ksqldbx/net"
)

func TestIntegrationKsqlDB_Pull(t *testing.T) {
	createStreamStmnt := `
	CREATE STREAM IF NOT EXISTS pull_test(
		d1 BOOLEAN,
		d2 VARCHAR,
		d3 INT,
		d4 BIGINT,
		d5 DOUBLE,
		d6 DECIMAL(5, 2)
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

	// drop stream
	defer func() {
		_, err = ksql.Exec(context.Background(), StmntSQL{KSQL: "DROP STREAM pull_test;"})
		require.Nil(t, err)
	}()

	type args struct {
		ctx context.Context
		q   QuerySQL
	}
	tests := []struct {
		name       string
		args       args
		initData   func(ctx context.Context, ksql *KsqlDB) error
		wantHeader Header
		wantRows   []Row
		wantErr    bool
	}{
		{
			name: "success running SELECT query",
			args: args{
				ctx: context.Background(),
				q: QuerySQL{
					SQL: "SELECT * FROM pull_test LIMIT 1;",
				},
			},
			initData: func(ctx context.Context, ksql *KsqlDB) error {
				// insert exactly one row of data
				stmnt := StmntSQL{
					KSQL: `
					INSERT INTO pull_test(
						d1,
						d2,
						d3,
						d4,
						d5,
						d6
					) VALUES (
						true,
						'varchar',
						1,
						10,
						1.5,
						10.50
					);`,
				}

				_, err := ksql.Exec(ctx, stmnt)
				return err
			},
			wantHeader: Header{
				ColumnNames: []string{"D1", "D2", "D3", "D4", "D5", "D6"},
				ColumnTypes: []string{"BOOLEAN", "STRING", "INTEGER", "BIGINT", "DOUBLE", "DECIMAL(5, 2)"},
			},
			wantRows: []Row{
				// by default data with data type INTEGER (or any integer for that matter)
				// will be decoded into type of float64
				{true, "varchar", float64(1), float64(10), 1.5, 10.50},
			},
			wantErr: false,
		},
		{
			name: "fail running invalid SELECT query",
			args: args{
				ctx: context.Background(),
				q: QuerySQL{
					SQL: "SELECT * blabla",
				},
			},
			initData: func(ctx context.Context, ksql *KsqlDB) error {
				return nil
			},
			wantHeader: Header{},
			wantRows:   nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		err := tt.initData(ctx, ksql)
		require.Nil(t, err)

		header, rows, err := ksql.Pull(tt.args.ctx, tt.args.q)
		assert.Equal(t, tt.wantHeader.ColumnNames, header.ColumnNames)
		assert.Equal(t, tt.wantHeader.ColumnTypes, header.ColumnTypes)
		assert.Equal(t, tt.wantRows, rows)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
