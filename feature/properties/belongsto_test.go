package properties

import (
	"io"
	"strings"
	"testing"
)

var has_belongs_to = strings.NewReader(`{"properties":{"wof:belongsto": [ 123, 456, 789 ] }}`)
var missing_belongs_to = strings.NewReader(`{"properties":{ }}`)

func TestHasBelongsTo(t *testing.T) {

	body, err := io.ReadAll(has_belongs_to)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	s := BelongsTo(body)

	if len(s) != 3 {
		t.Fatal("Invalid belongs_to value")
	}

}

func TestMissingBelongsto(t *testing.T) {

	body, err := io.ReadAll(missing_belongs_to)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	s := BelongsTo(body)

	if len(s) != 0 {
		t.Fatal("Invalid belongs_to (missing) value")
	}

}
