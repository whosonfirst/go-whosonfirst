package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_hierarchies = strings.NewReader(`{"properties":{"wof:hierarchy": [ {"foo_id": 123 }, { "bar_id": 456 } ]}}`)

var valid_hierarchies_2 = strings.NewReader(`{"properties":{"wof:hierarchy": [ {"foo_id": 123 }, { "bar_id": 456 }, { "hello": 1, "world": 2 } ]}}`)

var missing_hierarchies = strings.NewReader(`{"properties":{ }}`)

func TestValidHierarchies(t *testing.T) {

	body, err := io.ReadAll(valid_hierarchies)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	hierarchies := Hierarchies(body)

	if hierarchies == nil {
		t.Fatalf("Missing hierarchies")
	}

	if len(hierarchies) != 2 {
		t.Fatal("Invalid hierarchies, expected (2) results")
	}

	if hierarchies[0]["foo_id"] != 123 {
		t.Fatalf("Invalid hierarchies[0][\"foo_id\"] value")
	}

}

func TestMissingHierarchies(t *testing.T) {

	body, err := io.ReadAll(missing_hierarchies)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	hierarchies := Hierarchies(body)

	if hierarchies != nil {
		t.Fatalf("Hierarchies should be nil")
	}
}

func TestMergeHierarchies(t *testing.T) {

	f1, err := io.ReadAll(missing_hierarchies)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	f2, err := io.ReadAll(valid_hierarchies)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	f3, err := io.ReadAll(valid_hierarchies_2)

	if err != nil {
		t.Fatalf("Failed to read data (valid 2), %v", err)
	}

	hierarchies, err := MergeHierarchies(f1, f2, f3)

	if err != nil {
		t.Fatalf("Failed to merge hierarchies, %v", err)
	}

	count := len(hierarchies)
	expected := 3

	if count != expected {
		t.Fatalf("Unexpected hierarchy count after merging. Expected %d but got %d", expected, count)
	}
}
