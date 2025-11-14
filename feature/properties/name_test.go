package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_name = strings.NewReader(`{"properties":{"wof:name": "Test" }}`)

var missing_name = strings.NewReader(`{"properties":{ }}`)

func TestValidName(t *testing.T) {

	body, err := io.ReadAll(valid_name)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	name, err := Name(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if name != "Test" {
		t.Fatal("Invalid NAME (valid)")
	}

}

func TestMissingName(t *testing.T) {

	body, err := io.ReadAll(missing_name)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = Name(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
