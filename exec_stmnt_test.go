package ksqldbx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tamboto2000/ksqldbx/net"
)

func TestIntegrationKsqlDB_Exec(t *testing.T) {
	ksql, err := NewKsqlDB(net.Options{
		BaseUrl:   "http://localhost:8088",
		AllowHTTP: true,
	})

	require.Nil(t, err)

	type args struct {
		ctx   context.Context
		stmnt StmntSQL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success on executing any valid statement",
			args: args{
				ctx: context.Background(),
				stmnt: StmntSQL{
					KSQL: "LIST STREAMS;",
				},
			},
			wantErr: false,
		},
		{
			name: "fail on executing any invalid statements",
			args: args{
				ctx: context.Background(),
				stmnt: StmntSQL{
					KSQL: "INVALID STATEMENT;",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		res, err := ksql.Exec(tt.args.ctx, tt.args.stmnt)
		if tt.wantErr {
			assert.Zero(t, res)
			assert.NotNil(t, err)
		} else {
			assert.NotZero(t, res)
			assert.Nil(t, err)
		}
	}
}
