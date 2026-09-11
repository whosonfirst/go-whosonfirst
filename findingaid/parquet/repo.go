package parquet

import (
	"fmt"

	sfom_parquet "github.com/sfomuseum/go-parquet"
)

func GetRepo(uri string, id int64) (string, error) {

	rec, err := sfom_parquet.FindRecordById[*ParquetRecord, int64](uri, id)

	if err != nil {
		return "", fmt.Errorf("Failed to retrieve record, %w", err)
	}

	return rec.Repo, nil
}
