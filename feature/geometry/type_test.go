package geometry

import (
	"io"
	"strings"
	"testing"
)

var valid_type = strings.NewReader(`{"geometry":{"type": "Point" }}`)

var missing_type = strings.NewReader(`{"properties":{ }}`)

func TestValidType(t *testing.T) {

	body, err := io.ReadAll(valid_type)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	geom_type, err := Type(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if geom_type != "Point" {
		t.Fatal("Invalid TYPE (valid)")
	}

}

func TestMissingType(t *testing.T) {

	body, err := io.ReadAll(missing_type)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = Type(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
