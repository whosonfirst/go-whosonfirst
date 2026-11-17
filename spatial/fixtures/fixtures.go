// Package fixtures provides fixtures for testing spatial queries.
package fixtures

import (
	"embed"
)

//go:embed *.geojson
var FS embed.FS
