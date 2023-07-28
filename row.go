package ksqldbx

import (
	"fmt"

	"github.com/tamboto2000/ksqldbx/internal/scanner"
)

type Row []any

func (r Row) Scan(dest ...any) error {
	if len(r) != len(dest) {
		return fmt.Errorf("scanning %d values into %d destinations", len(r), len(dest))
	}

	for i, v := range r {
		if err := scanner.Scan(dest[i], v); err != nil {
			return fmt.Errorf("error on scanning destination at index %d: %s", i, err.Error())
		}
	}

	return nil
}
