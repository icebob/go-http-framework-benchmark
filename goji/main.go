package main

import (
	"fmt"
	"io"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	addr := "127.0.0.1:8006"
	fmt.Println("Listening on ", addr)

	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/"), hello)

	http.ListenAndServe(addr, mux)
}
