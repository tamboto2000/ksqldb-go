package ksqldbx

import "context"

// Describe get details on a TABLE or STREAM by sourceName.
// See https://docs.ksqldb.io/en/latest/developer-guide/ksqldb-reference/describe/
// for more info
func (ksql *KsqlDB) Describe(ctx context.Context, sourceName string) (SourceDesc, error) {
	var srcDesc SourceDesc
	res, err := ksql.Exec(ctx, StmntSQL{
		KSQL: "DESCRIBE ${src_name};",
		Variables: Variables{
			"src_name": sourceName,
		},
	})

	if err != nil {
		return srcDesc, err
	}

	return *res[0].SourceDesc, nil
}
