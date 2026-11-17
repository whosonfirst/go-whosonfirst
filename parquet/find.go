// Package parquet provides helpers for reading Who's On First records from Parquet files.
package parquet

import (
	sfom_parquet "github.com/sfomuseum/go-parquet"
)

func FindRecord(uri string, id int64) (*Record, error) {
	return sfom_parquet.FindRecordById[*Record, int64](uri, id)
}
