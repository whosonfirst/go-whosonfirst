//go:build wasmjs
package main

import (
	"log"
	"syscall/js"

	"github.com/whosonfirst/go-whosonfirst/v4/findingaid/parquet/wasm"
)

func main() {

	get_repo_func := wasm.GetRepoFunc()
	defer get_repo_func.Release()

	js.Global().Set("parquet_get_repo", get_repo_func)

	c := make(chan struct{}, 0)

	log.Println("parquet_get_repo function initialized")
	<-c
}

