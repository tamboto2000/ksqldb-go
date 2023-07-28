package ksqldbx

type Header struct {
	QueryID     string   `json:"queryId"`
	ColumnNames []string `json:"columnNames"`
	ColumnTypes []string `json:"columnTypes"`
}
