package ksqldbx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamboto2000/ksqldbx/net"
)

func TestIntegrationKsqlDB_Describe(t *testing.T) {
	createStreamStmnt := `CREATE STREAM IF NOT EXISTS describe_stream_test (
		d1 VARCHAR
	) WITH (
		kafka_topic='describe_stream_test', 
		value_format='protobuf',
		partitions=1
	);`

	createTableSmnt := `CREATE TABLE describe_table_test (
		d1 VARCHAR PRIMARY KEY,
		d2 VARCHAR
	) WITH (
		kafka_topic='describe_table_test', 
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
		_, err = ksql.Exec(context.Background(), StmntSQL{KSQL: "DROP STREAM describe_stream_test;"})
		require.Nil(t, err)
	}()

	// create a table
	_, err = ksql.Exec(context.Background(), StmntSQL{
		KSQL: createTableSmnt,
	})

	// drop table
	defer func() {
		_, err = ksql.Exec(context.Background(), StmntSQL{KSQL: "DROP TABLE describe_table_test;"})
		require.Nil(t, err)
	}()

	require.Nil(t, err)

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantSrc SourceDesc
		wantErr bool
	}{
		{
			name: "success running DESCRIBE query",
			args: args{
				ctx:  context.Background(),
				name: "describe_stream_test",
			},
			wantSrc: SourceDesc{
				Name:         "DESCRIBE_STREAM_TEST",
				ReadQueries:  []Query{},
				WriteQueries: []Query{},
				Fields: []Field{
					{
						Name: "D1",
						Schema: SchemaNode{
							Schema: Schema{
								Type:   "STRING",
								Fields: nil,
							},
						},
					},
				},
				Type:        "STREAM",
				Extended:    false,
				KeyFormat:   "KAFKA",
				ValueFormat: "PROTOBUF",
				Topic:       "describe_stream_test",
				Partitions:  1,
				Replication: 1,
			},
			wantErr: false,
		},
		{
			name: "fail running DESCRIBE query on non existing STREAM",
			args: args{
				ctx:  context.Background(),
				name: "abcd",
			},
			wantSrc: SourceDesc{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		src, err := ksql.Describe(tt.args.ctx, tt.args.name)
		assert.Equal(t, tt.wantSrc, src)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
