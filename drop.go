package ksqldbx

import "context"

func (ksql *KsqlDB) DropStream(ctx context.Context, name string) error {
	stmnt := StmntSQL{
		KSQL: "DROP STREAM ${stream_name};",
		Variables: Variables{
			"stream_name": name,
		},
	}

	_, err := ksql.Exec(ctx, stmnt)
	return err
}

func (ksql *KsqlDB) DropTable(ctx context.Context, name string) error {
	stmnt := StmntSQL{
		KSQL: "DROP TABLE ${table_name};",
		Variables: Variables{
			"table_name": name,
		},
	}

	_, err := ksql.Exec(ctx, stmnt)
	return err
}
