package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_repo = strings.NewReader(`{"properties":{"wof:repo": "whosonfirst-data-test" }}`)

var missing_repo = strings.NewReader(`{"properties":{ }}`)

func TestValidRepo(t *testing.T) {

	body, err := io.ReadAll(valid_repo)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	repo, err := Repo(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if repo != "whosonfirst-data-test" {
		t.Fatal("Invalid REPO (valid)")
	}

}

func TestMissingRepo(t *testing.T) {

	body, err := io.ReadAll(missing_repo)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = Repo(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
