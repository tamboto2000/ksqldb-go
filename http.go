package ksqldbx

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

func (ksql *KsqlDB) queryStreamRequest(ctx context.Context, q QuerySQL) (*http.Response, error) {
	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	return ksql.postRequest(ctx, endpStreamQuery, b)
}

func (ksql *KsqlDB) queryCloseRequest(ctx context.Context, qId string) (*http.Response, error) {
	q := QuerySQL{QueryID: qId}
	b, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	return ksql.postRequest(ctx, endpCloseQuery, b)
}

func (ksql *KsqlDB) ksqlRequest(ctx context.Context, stmnt StmntSQL) (*http.Response, error) {
	b, err := json.Marshal(stmnt)
	if err != nil {
		return nil, err
	}

	return ksql.postRequest(ctx, endpKsql, b)
}

func (ksql *KsqlDB) postRequest(ctx context.Context, endp string, body []byte) (*http.Response, error) {
	reader := bytes.NewReader(body)
	return ksql.cl.Post(ctx, endp, "", reader)
}
