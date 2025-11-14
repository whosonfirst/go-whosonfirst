package alt

import (
	"io"
	"strings"
	"testing"
)

var valid_alt = strings.NewReader(`{"properties":{"src:alt_label": "naturaleaarth" }}`)
var valid_alt_old = strings.NewReader(`{"properties":{"wof:alt_label": "zetashapes" }}`)

var missing_alt = strings.NewReader(`{"properties":{ }}`)

func TestValidIsAlt(t *testing.T) {

	valid := map[string]io.Reader{
		"src:alt_label": valid_alt,
		"wof:alt_label": valid_alt_old,
	}

	for path, r := range valid {

		body, err := io.ReadAll(r)

		if err != nil {
			t.Fatalf("Failed to read data (%s), %v", path, err)
		}

		is_alt := IsAlt(body)

		if !is_alt {
			t.Fatalf("Expected result (%s) failed", path)
		}
	}
}

func TestNotIsAlt(t *testing.T) {

	body, err := io.ReadAll(missing_alt)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	is_alt := IsAlt(body)

	if is_alt {
		t.Fatalf("Expected result (missing) failed")
	}
}
