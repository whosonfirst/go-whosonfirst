package properties

import (
	"github.com/sfomuseum/go-edtf"
	"io"
	"strings"
	"testing"
)

var has_inception = strings.NewReader(`{"properties":{"edtf:inception": "2021-10-07" }}`)
var missing_inception = strings.NewReader(`{"properties":{ }}`)

var has_cessation = strings.NewReader(`{"properties":{"edtf:cessation": "2021-10-07" }}`)
var missing_cessation = strings.NewReader(`{"properties":{ }}`)

var has_deprecated = strings.NewReader(`{"properties":{"edtf:deprecated": "2021-10-07" }}`)
var missing_deprecated = strings.NewReader(`{"properties":{ }}`)

func TestHasInception(t *testing.T) {

	body, err := io.ReadAll(has_inception)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	edtf_str := Inception(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if edtf_str != "2021-10-07" {
		t.Fatalf("Invalid inception value '%s'", edtf_str)
	}

}

func TestMissingInception(t *testing.T) {

	body, err := io.ReadAll(missing_inception)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	edtf_str := Inception(body)

	if edtf_str != edtf.UNKNOWN {
		t.Fatalf("Invalid inception (missing) value '%s'", edtf_str)
	}

}

func TestHasCessation(t *testing.T) {

	body, err := io.ReadAll(has_cessation)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	edtf_str := Cessation(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if edtf_str != "2021-10-07" {
		t.Fatalf("Invalid cessation value '%s'", edtf_str)
	}

}

func TestMissingCessation(t *testing.T) {

	body, err := io.ReadAll(missing_cessation)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	edtf_str := Cessation(body)

	if edtf_str != edtf.UNKNOWN {
		t.Fatalf("Invalid cessation (missing) value, '%s'", edtf_str)
	}

}

func TestHasDeprecated(t *testing.T) {

	body, err := io.ReadAll(has_deprecated)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	edtf_str := Deprecated(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if edtf_str != "2021-10-07" {
		t.Fatalf("Invalid deprecated value '%s'", edtf_str)
	}

}

func TestMissingDeprecated(t *testing.T) {

	body, err := io.ReadAll(missing_deprecated)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	edtf_str := Deprecated(body)

	if edtf_str != "" {
		t.Fatalf("Invalid deprecated (missing) value, '%s'", edtf_str)
	}

}
