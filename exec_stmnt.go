package ksqldbx

import (
	"context"
	"encoding/json"
	"net/http"
)

type StmntSQL struct {
	KSQL         string     `json:"ksql"`
	Properties   Properties `json:"streamsProperties"`
	Variables    Variables  `json:"sessionVariables"`
	CmdSeqNumber int64      `json:"commandSequenceNumber"`
}

type Warning struct {
	Message string `json:"message"`
}

type CmdStatus struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	CmdSeqNumber int64  `json:"commandSequenceNumber"`
}

type Table struct {
	Name       string `json:"name"`
	Topic      string `json:"topic"`
	Format     string `json:"format"`
	Type       string `json:"type"`
	IsWindowed bool   `json:"isWindowed"`
}

type Query struct {
	QueryString string `json:"queryString"`
	Sinks       string `json:"sinks"`
	ID          string `json:"id"`
}

type SchemaNode struct {
	Schema
	MemberSchema Schema `json:"memberSchema"`
}

type Schema struct {
	Type   string  `json:"type"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name   string     `json:"name"`
	Schema SchemaNode `json:"schema"`
}

type SourceDesc struct {
	Name         string  `json:"name"`
	ReadQueries  []Query `json:"readQueries"`
	WriteQueries []Query `json:"writeQueries"`
	Fields       []Field `json:"fields"`
	Type         string  `json:"type"`
	Key          string  `json:"key"`
	Timestamp    string  `json:"timestamp"`
	Format       string  `json:"format"`
	Topic        string  `json:"topic"`
	KeyFormat    string  `json:"keyFormat"`
	ValueFormat  string  `json:"valueFormat"`
	IsWindowed   bool    `json:"false"`
	Extended     bool    `json:"extended"`
	Statistics   string  `json:"statistics"`
	ErrorStats   string  `json:"errorStats"`
	Replication  int     `json:"replication"`
	Partitions   int     `json:"partitions"`
}

type QueryDesc struct {
	StmntText            string            `json:"statementText"`
	Fields               []Field           `json:"fields"`
	Sources              []string          `json:"sources"`
	Sinks                []string          `json:"sinks"`
	ExecutionPlan        string            `json:"executionPlan"`
	Topology             string            `json:"topology"`
	OverriddenProperties map[string]string `json:"overriddenProperties"`
}

type Response struct {
	StmntText  string      `json:"statementText"`
	Warnings   []Warning   `json:"warnings"`
	CmdID      string      `json:"commandId,omitempty"`
	CmdStatus  *CmdStatus  `json:"commandStatus,omitempty"`
	Streams    []Stream    `json:"streams,omitempty"`
	Tables     []Table     `json:"tables,omitempty"`
	Queries    []Query     `json:"queries,omitempty"`
	Properties *Properties `json:"properties,omitempty"`
	SourceDesc *SourceDesc `json:"sourceDescription,omitempty"`
	QueryDesc  *QueryDesc  `json:"queryDescription"`
}

// Exec execute KSQL statements.
// Refer to https://docs.ksqldb.io/en/latest/developer-guide/ksqldb-rest-api/ksql-endpoint/
// for more info
func (ksql *KsqlDB) Exec(ctx context.Context, stmnt StmntSQL) ([]Response, error) {
	var kres []Response
	res, err := ksql.ksqlRequest(ctx, stmnt)
	if err != nil {
		return kres, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var err Error
		if err := json.NewDecoder(res.Body).Decode(&err); err != nil {
			return kres, err
		}

		return kres, err
	}

	if err := json.NewDecoder(res.Body).Decode(&kres); err != nil {
		return kres, err
	}

	return kres, nil
}
