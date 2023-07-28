package ksqldbx

import (
	"context"
	"encoding/json"
	"net/http"
)

func (ksql *KsqlDB) CloseQuery(ctx context.Context, qId string) error {
	res, err := ksql.queryCloseRequest(ctx, qId)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var err Error
		if err := json.NewDecoder(res.Body).Decode(&err); err != nil {
			return err
		}

		return err
	}

	return nil
}
