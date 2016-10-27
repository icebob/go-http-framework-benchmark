package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// request handler in fasthttp style, i.e. just plain function.
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Hello World!")
}

func main() {
	addr := "127.0.0.1:8003"
	fmt.Println("Listening on ", addr)

	fasthttp.ListenAndServe(addr, fastHTTPHandler)
}
