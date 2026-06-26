package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_parent_id = strings.NewReader(`{"properties":{"wof:parent_id": 1234 }}`)

var invalid_parent_id = strings.NewReader(`{"properties":{"wof:parent_id": -5 }}`)

var missing_parent_id = strings.NewReader(`{"properties":{ }}`)

func TestValidParentId(t *testing.T) {

	body, err := io.ReadAll(valid_parent_id)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	id, err := ParentId(body)

	if err != nil {
		t.Fatalf("Expected data (%d, valid) failed", id)
	}

	if id != 1234 {
		t.Fatal("Invalid ID (valid)")
	}

}

func TestInValidParentId(t *testing.T) {

	body, err := io.ReadAll(invalid_parent_id)

	if err != nil {
		t.Fatalf("Failed to read data (invalid), %v", err)
	}

	id, err := ParentId(body)

	if err == nil {
		t.Fatalf("Expected data (%d, invalid) to fail", id)
	}
}

func TestMissingParentId(t *testing.T) {

	body, err := io.ReadAll(missing_parent_id)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = ParentId(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
