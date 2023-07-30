package ksqldbx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamboto2000/ksqldbx/net"
)

func TestIntegrationKsqlDB_DropStream(t *testing.T) {
	createStreamStmnt := `
	CREATE STREAM IF NOT EXISTS drop_stream_test(
		d1 VARCHAR
	) WITH (
		kafka_topic='drop_stream_test', 
		value_format='protobuf',
		partitions=1
	);
	`

	ksql, err := NewKsqlDB(net.Options{
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	})

	require.Nil(t, err)

	// create a stream
	_, err = ksql.Exec(context.Background(), StmntSQL{KSQL: createStreamStmnt})
	require.Nil(t, err)

	// drop stream
	defer func() {
		_, err := ksql.Exec(context.Background(), StmntSQL{KSQL: "DROP STREAM IF EXISTS drop_stream_test;"})
		require.Nil(t, err)
	}()

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success dropping a STREAM",
			args: args{
				ctx:  context.Background(),
				name: "drop_stream_test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		err := ksql.DropStream(tt.args.ctx, tt.args.name)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestIntegrationKsqlDB_DropTable(t *testing.T) {
	createTableStmnt := `
	CREATE TABLE IF NOT EXISTS drop_table_test(
		d1 VARCHAR PRIMARY KEY,
		d2 VARCHAR
	) WITH (
		kafka_topic='drop_table_test', 
		value_format='protobuf',
		partitions=1
	);
	`

	ksql, err := NewKsqlDB(net.Options{
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	})

	require.Nil(t, err)

	// create a table
	_, err = ksql.Exec(context.Background(), StmntSQL{KSQL: createTableStmnt})
	require.Nil(t, err)

	// drop table
	defer func() {
		_, err := ksql.Exec(context.Background(), StmntSQL{KSQL: "DROP TABLE IF EXISTS drop_table_test;"})
		require.Nil(t, err)
	}()

	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success dropping a TABLE",
			args: args{
				ctx:  context.Background(),
				name: "drop_table_test",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		err := ksql.DropTable(tt.args.ctx, tt.args.name)
		if tt.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
