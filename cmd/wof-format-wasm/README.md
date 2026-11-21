## wof-format-wasm

Registers the Who's On First GeoJSON formatter as a JavaScript function.

This is a WebAssembly (`GOOS=js GOARCH=wasm`, `-tags wasmjs`) binary and does not accept command-line flags (`-h`). It registers the following JavaScript global function(s):

- `wof_format`
