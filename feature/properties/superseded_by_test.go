package properties

import (
	"io"
	"strings"
	"testing"
)

var has_superseded_by = strings.NewReader(`{"properties":{"wof:superseded_by": [ 123, 456, 789 ] }}`)
var missing_superseded_by = strings.NewReader(`{"properties":{ }}`)

func TestHasSupersededBy(t *testing.T) {

	body, err := io.ReadAll(has_superseded_by)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	s := SupersededBy(body)

	if len(s) != 3 {
		t.Fatal("Invalid superseded_by value")
	}

}

func TestMissingSupersededBy(t *testing.T) {

	body, err := io.ReadAll(missing_superseded_by)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	s := SupersededBy(body)

	if len(s) != 0 {
		t.Fatal("Invalid superseded_by (missing) value")
	}

}
