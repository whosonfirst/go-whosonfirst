// Package microhoods provides neighborhood fixtures for spatial tests.
package microhoods

import (
	"embed"
)

//go:embed *.geojson
var FS embed.FS
