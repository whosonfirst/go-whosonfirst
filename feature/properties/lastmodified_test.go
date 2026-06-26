package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_lastmod = strings.NewReader(`{"properties":{"wof:lastmodified": 1635460530 }}`)

var missing_lastmod = strings.NewReader(`{"properties":{ }}`)

func TestValidLastModified(t *testing.T) {

	body, err := io.ReadAll(valid_lastmod)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	lastmod := LastModified(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if lastmod != 1635460530 {
		t.Fatal("Invalid LASTMOD (valid)")
	}

}

func TestMissingLastmodified(t *testing.T) {

	body, err := io.ReadAll(missing_lastmod)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	lastmod := LastModified(body)

	if err != nil {
		t.Fatalf("Expect data (missing) failed")
	}

	if lastmod != -1 {
		t.Fatal("Invalid LASTMOD (missing)")
	}
}
