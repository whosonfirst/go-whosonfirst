package sort

import (
	"github.com/whosonfirst/go-whosonfirst/v4/spr"
)

type testableStandardPlacesResults struct {
	spr.StandardPlacesResults
	results []spr.StandardPlacesResult
}

func (r *testableStandardPlacesResults) Results() []spr.StandardPlacesResult {
	return r.results
}
