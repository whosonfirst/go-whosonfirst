package properties

import (
	_ "fmt"
	"io"
	"strings"
	"testing"
)

var valid_country_1 = strings.NewReader(`{"properties":{"wof:country": "CA" }}`)

var valid_country_2 = strings.NewReader(`{"properties":{"wof:country": "CA" }}`)

var valid_country_3 = strings.NewReader(`{"properties":{"wof:country": "US" }}`)

var missing_country = strings.NewReader(`{"properties":{ }}`)

func TestValidCountry(t *testing.T) {

	err := resetCountryReaders()

	if err != nil {
		t.Fatalf("Failed to reset readers, %v", err)
	}

	body, err := io.ReadAll(valid_country_1)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	country := Country(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if country != "CA" {
		t.Fatal("Invalid COUNTRY (valid)")
	}

}

func TestMissingCountry(t *testing.T) {

	err := resetCountryReaders()

	if err != nil {
		t.Fatalf("Failed to reset readers, %v", err)
	}

	body, err := io.ReadAll(missing_country)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	country := Country(body)

	if err != nil {
		t.Fatalf("Expect data (missing) failed")
	}

	if country != COUNTRY_UNKNOWN {
		t.Fatal("Invalid COUNTRY (missing)")
	}
}

func TestMergeCountries(t *testing.T) {

	err := resetCountryReaders()

	if err != nil {
		t.Fatalf("Failed to reset readers, %v", err)
	}

	f1, err := io.ReadAll(valid_country_1)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	f2, err := io.ReadAll(valid_country_2)

	if err != nil {
		t.Fatalf("Failed to read data (valid 2), %v", err)
	}

	f3, err := io.ReadAll(valid_country_3)

	if err != nil {
		t.Fatalf("Failed to read data (valid 3), %v", err)
	}

	f4, err := io.ReadAll(missing_country)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	missing := MergeCountries(f4)

	if missing != COUNTRY_UNKNOWN {
		t.Fatalf("Expected XY for missing country code, but got %s", missing)
	}

	unknown := MergeCountries(f1, f2, f3, f4)

	if unknown != COUNTRY_COMPLICATED {
		t.Fatalf("Expected XX for multiple country codes, but got %s", unknown)
	}

	ca := MergeCountries(f1, f2)

	if ca != "CA" {
		t.Fatalf("Expected CA for Canada, but got %s", ca)
	}

}

func resetCountryReaders() error {

	readers := []io.ReadSeeker{
		valid_country_1,
		valid_country_2,
		valid_country_3,
		missing_country,
	}

	return resetReaders(readers...)
}

func resetReaders(readers ...io.ReadSeeker) error {

	for _, r := range readers {

		_, err := r.Seek(0, 0)

		if err != nil {
			return err
		}
	}

	return nil
}
