package ksqldbx

import (
	"github.com/tamboto2000/ksqldbx/net"
)

type KsqlDB struct {
	cl net.HTTPClient
}

// New instatiate KsqlDB
func NewKsqlDB(opts net.Options) (*KsqlDB, error) {
	cl, err := net.NewHTTPClient(opts, nil)
	if err != nil {
		return nil, err
	}

	ksql := &KsqlDB{
		cl: cl,
	}

	return ksql, nil
}
