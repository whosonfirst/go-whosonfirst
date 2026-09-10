package parquet

import (
	"fmt"

	sfom_parquet "github.com/sfomuseum/go-parquet"
	"github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func FindRecord(uri string, id int64, args ...*uri.URIArgs) (*Record, error) {

	records, err := FindRecords(uri, id)

	if err != nil {
		return nil, err
	}

	possible := make([]*Record, 0)

	switch len(args) == 1 && args[0].IsAlternate {
	case true:

		uri_args := args[0]
		alt_geom := uri_args.AltGeom
		alt_label, err := alt_geom.String()

		if err != nil {
			return nil, fmt.Errorf("Failed to derive label for alt geometry, %w", err)
		}

		for _, r := range records {

			if r.AltLabel == alt_label {
				possible = append(possible, r)
			}
		}

	default:

		for _, r := range records {

			if r.AltLabel == "" {
				possible = append(possible, r)
			}
		}
	}

	switch len(possible) {
	case 1:
		return possible[0], nil
	case 0:
		return nil, fmt.Errorf("Not found")
	default:
		return nil, fmt.Errorf("Multiple matches")
	}

}

func FindRecords(uri string, id int64) ([]*Record, error) {
	return sfom_parquet.FindRecordByIdAll[*Record, int64](uri, id)
}
