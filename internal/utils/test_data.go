package utils

import "fmt"

// CreateStreamTestStmnt is a function to create
// STREAM with name. The sql statement produced
// by this function can be used for creating
// test stream for integration testing
func CreateStreamTestStmnt(name string) string {
	stmnt := `
		CREATE STREAM IF NOT EXISTS %s(
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
		);
	`

	return fmt.Sprintf(stmnt, name)
}

func SelectStreamTestQuery() string {
	q := `
		SELECT
			d1,
			d2,
			d3,
			d4,
			d5,
			d6,
			d7
		FROM test_stream;		
	`

	return q
}
