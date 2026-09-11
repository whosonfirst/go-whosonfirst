//go:build wasmjs
package wasm

import (
	"fmt"
	"syscall/js"
	"strconv"
	
	"github.com/whosonfirst/go-whosonfirst/v4/findingaid/parquet"
)

func GetRepoFunc() js.Func {

	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		parquet_uri := args[0].String()
		str_id := args[1].String()

		handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			resolve := args[0]
			reject := args[1]

			go func() {

				id, err := strconv.ParseInt(str_id, 10, 64)
				
				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to parse ID, %w", err))
					return
				}

				repo, err := parquet.GetRepo(parquet_uri, id)

				if err != nil {
					reject.Invoke(fmt.Sprintf("Failed to derive repo, %w", err))
					return
				}

				resolve.Invoke(string(repo))
			}()

			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

