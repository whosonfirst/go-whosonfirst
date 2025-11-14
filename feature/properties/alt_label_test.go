package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_label = strings.NewReader(`{"properties":{"src:alt_label": "quattroshapes" }}`)

var missing_label = strings.NewReader(`{"properties":{ }}`)

func TestValidAltLabel(t *testing.T) {

	body, err := io.ReadAll(valid_label)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	alt_label, err := AltLabel(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if alt_label != "quattroshapes" {
		t.Fatal("Invalid LABEL (valid)")
	}

}

func TestMissingLabel(t *testing.T) {

	body, err := io.ReadAll(missing_label)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	alt_label, err := AltLabel(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if alt_label != "" {
		t.Fatalf("Expect data (missing) to fail")
	}
}
