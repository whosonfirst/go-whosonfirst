// Package static provides static assets for the wof editor.
package static

import (
	"embed"
)

//go:embed css/* javascript/* wasm/* *.html
var FS embed.FS
