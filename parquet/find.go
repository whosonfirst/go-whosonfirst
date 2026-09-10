package parquet

import (
	"fmt"

	sfom_parquet "github.com/sfomuseum/go-parquet"
	"github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func FindRecord(uri string, id int64, args ...*uri.URIArgs) (*Record, error) {

	switch len(args) {
	case 1:

		uri_args := args[0]

		switch uri_args.IsAlternate {
		case true:

			records, err := FindRecords(uri, id)

			if err != nil {
				return nil, err
			}

			for range records {
				// TEST ALT LABEL(S) HERE
			}

			return nil, fmt.Errorf("Not found")

		default:
			return sfom_parquet.FindRecordById[*Record, int64](uri, id)
		}

	default:
		return sfom_parquet.FindRecordById[*Record, int64](uri, id)
	}
}

func FindRecords(uri string, id int64) ([]*Record, error) {
	return sfom_parquet.FindRecordByIdAll[*Record, int64](uri, id)
}
