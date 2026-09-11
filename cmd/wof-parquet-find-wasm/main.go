//go:build wasmjs
package main

import (
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst/v4/parquet/wasm"
)

func main() {

	find_func := wasm.FindRecordFunc()
	defer find_func.Release()

	js.Global().Set("parquet_find_record", find_func)

	c := make(chan struct{}, 0)

	log.Println("parquet_find_record function initialized")
	<-c
}

