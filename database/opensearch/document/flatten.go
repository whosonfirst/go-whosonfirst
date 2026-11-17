package document

import (
	"context"
	_ "log"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Flatten merges the fields of the input JSON object's immediate child objects into a single JSON object.
func Flatten(ctx context.Context, body []byte) ([]byte, error) {

	var flattened []byte
	var err error

	rsp := gjson.ParseBytes(body)

	for _, details := range rsp.Map() {

		for k, v := range details.Map() {

			flattened, err = sjson.SetBytes(flattened, k, v.Value())

			if err != nil {
				return nil, err
			}
		}

	}

	return flattened, nil
}
