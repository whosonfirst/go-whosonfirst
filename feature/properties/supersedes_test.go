package properties

import (
	"io"
	"strings"
	"testing"
)

var has_supersedes = strings.NewReader(`{"properties":{"wof:supersedes": [ 123, 456, 789 ] }}`)
var missing_supersedes = strings.NewReader(`{"properties":{ }}`)

func TestHasSupersedes(t *testing.T) {

	body, err := io.ReadAll(has_supersedes)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	s := Supersedes(body)

	if len(s) != 3 {
		t.Fatal("Invalid supersedes value")
	}

}

func TestMissingSupersedes(t *testing.T) {

	body, err := io.ReadAll(missing_supersedes)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	s := Supersedes(body)

	if len(s) != 0 {
		t.Fatal("Invalid supersedes (missing) value")
	}

}
