package iterate

import (
	"strings"
	"testing"

	"github.com/whosonfirst/go-ioutil"
)

func TestReadAllAndClose(t *testing.T) {

	expected := "Hello world"

	sr := strings.NewReader(expected)

	rcs, err := ioutil.NewReadSeekCloser(sr)

	if err != nil {
		t.Fatalf("Failed to create ReadSeekCloser, %v", err)
	}

	r := &Record{
		Path: "local",
		Body: rcs,
	}

	body, err := r.ReadAllAndClose()

	if err != nil {
		t.Fatalf("Failed to read and close, %v", err)
	}

	if string(body) != expected {
		t.Fatalf("Unexpected output, %s", string(body))
	}
}
