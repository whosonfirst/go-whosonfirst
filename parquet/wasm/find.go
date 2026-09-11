//go:build wasmjs

package wasm

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
	"github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func FindRecordFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		parquet_uri := args[0].String()
		wof_uri := args[1].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				id, uri_args, err := uri.ParseURI(wof_uri)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to parse URI, %w", err))
					return
				}

				rec, err := parquet.FindRecord(parquet_uri, id, uri_args)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to retrieve record, %w", err))
					return
				}

				f, err := rec.AsGeoJSON()

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to derive GeoJSON for record, %w", err))
					return
				}

				enc_f, err := json.Marshal(f)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to marshal GeoJSON, %w", err))
					return
				}

				resolve.Invoke(string(enc_f))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
