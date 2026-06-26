package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_concordances = strings.NewReader(`{"properties":{"wof:concordances": {"foo": "bar" }}}`)

var missing_concordances = strings.NewReader(`{"properties":{ }}`)

func TestValidConcordances(t *testing.T) {

	body, err := io.ReadAll(valid_concordances)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	concordances := Concordances(body)

	if concordances == nil {
		t.Fatalf("Missing concordances")
	}

	if concordances["foo"] != "bar" {
		t.Fatal("Invalid concordance")
	}

}

func TestMissingConcordances(t *testing.T) {

	body, err := io.ReadAll(missing_concordances)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	concordances := Concordances(body)

	_, ok := concordances["foo"]

	if ok {
		t.Fatalf("Expected data (missing) to fail")
	}
}
