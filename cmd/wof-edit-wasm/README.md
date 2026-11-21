## wof-edit-wasm

Registers JavaScript functions for preparing, formatting, and validating records and looking up placetypes in the wof editor.

This is a WebAssembly (`GOOS=js GOARCH=wasm`, `-tags wasmjs`) binary and does not accept command-line flags (`-h`). It registers the following JavaScript global function(s):

- `wof_prepare_feature`
- `wof_placetypes`
- `wof_placetypes_is_valid`
- `wof_placetypes_children`
- `wof_placetypes_descendants`
- `wof_placetypes_ancestors`
- `wof_format`
- `wof_validate`
