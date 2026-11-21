## wof-placetypes-wasm

Registers JavaScript functions for listing and validating placetypes and following their relationships.

This is a WebAssembly (`GOOS=js GOARCH=wasm`, `-tags wasmjs`) binary and does not accept command-line flags (`-h`). It registers the following JavaScript global function(s):

- `wof_placetypes`
- `wof_placetypes_is_valid`
- `wof_placetypes_children`
- `wof_placetypes_descendants`
- `wof_placetypes_ancestors`
