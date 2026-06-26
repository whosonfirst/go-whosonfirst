package properties

import (
	"io"
	"strings"
	"testing"
)

var current_known = strings.NewReader(`{"properties": {"mz:is_current": 1 }}`)

var current_unknown = strings.NewReader(`{"properties": { }}`)

var ceased = strings.NewReader(`{"properties": {"edtf:cessation": "2021-01-01" }}`)

var deprecated = strings.NewReader(`{"properties": {"edtf:deprecated": "2021-01-01" }}`)

var is_superseding = strings.NewReader(`{"properties": {"wof:supersedes": [ 1234 ] }}`)

var is_superseded_by = strings.NewReader(`{"properties": {"wof:superseded_by": [ 1234 ] }}`)

func TestIsCurrent(t *testing.T) {

	body_known, err := io.ReadAll(current_known)

	if err != nil {
		t.Fatalf("Failed to read data (known), %v", err)
	}

	fl_known, err := IsCurrent(body_known)

	if err != nil {
		t.Fatalf("Failed to determined if current (known), %v", err)
	}

	if !fl_known.IsTrue() {
		t.Fatalf("Is true test failed (known)")
	}

	if !fl_known.IsKnown() {
		t.Fatalf("Is known test failed (known)")
	}

	body_unknown, err := io.ReadAll(current_unknown)

	if err != nil {
		t.Fatalf("Failed to read data (unknown), %v", err)
	}

	fl_unknown, err := IsCurrent(body_unknown)

	if err != nil {
		t.Fatalf("Failed to determined if current (unknown), %v", err)
	}

	if fl_unknown.IsKnown() {
		t.Fatalf("Is known test failed (unknown)")
	}
}

func TestIsDeprecated(t *testing.T) {

	body, err := io.ReadAll(deprecated)

	if err != nil {
		t.Fatalf("Failed to read data, %v", err)
	}

	fl, err := IsDeprecated(body)

	if err != nil {
		t.Fatalf("Failed to determined if deprecated, %v", err)
	}

	if !fl.IsTrue() {
		t.Fatalf("Is true test failed")
	}

	if !fl.IsKnown() {
		t.Fatalf("Is known test failed")
	}

}

func TestIsCeased(t *testing.T) {

	body, err := io.ReadAll(ceased)

	if err != nil {
		t.Fatalf("Failed to read data, %v", err)
	}

	fl, err := IsCeased(body)

	if err != nil {
		t.Fatalf("Failed to determined if ceased, %v", err)
	}

	if !fl.IsTrue() {
		t.Fatalf("Is true test failed")
	}

	if !fl.IsKnown() {
		t.Fatalf("Is known test failed")
	}

}

func TestIsSuperseded(t *testing.T) {

	body, err := io.ReadAll(is_superseded_by)

	if err != nil {
		t.Fatalf("Failed to read data, %v", err)
	}

	fl, err := IsSuperseded(body)

	if err != nil {
		t.Fatalf("Failed to determined if superseded, %v", err)
	}

	if !fl.IsTrue() {
		t.Fatalf("Is true test failed")
	}

	if !fl.IsKnown() {
		t.Fatalf("Is known test failed")
	}
}

func TestIsSuperseding(t *testing.T) {

	body, err := io.ReadAll(is_superseding)

	if err != nil {
		t.Fatalf("Failed to read data, %v", err)
	}

	fl, err := IsSuperseding(body)

	if err != nil {
		t.Fatalf("Failed to determined if superseding, %v", err)
	}

	if !fl.IsTrue() {
		t.Fatalf("Is true test failed")
	}

	if !fl.IsKnown() {
		t.Fatalf("Is known test failed")
	}
}
