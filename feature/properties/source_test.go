package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_source_geom = strings.NewReader(`{"properties":{"src:geom": "example" }}`)
var valid_source_alt_label = strings.NewReader(`{"properties":{"src:alt_label": "example" }}`)

var missing_source = strings.NewReader(`{"properties":{ }}`)

func TestValidSourceGeom(t *testing.T) {

	body, err := io.ReadAll(valid_source_geom)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	source, err := Source(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if source != "example" {
		t.Fatal("Invalid source (valid)")
	}

}

func TestValidSourceAltLabel(t *testing.T) {

	body, err := io.ReadAll(valid_source_alt_label)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	source, err := Source(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if source != "example" {
		t.Fatalf("Invalid source (valid), '%s'", source)
	}

}

func TestMissingSource(t *testing.T) {

	body, err := io.ReadAll(missing_source)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = Source(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
