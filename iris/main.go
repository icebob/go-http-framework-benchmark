package main

import (
	"fmt"

	"github.com/kataras/iris"
)

func main() {
	addr := "127.0.0.1:8008"
	fmt.Println("Listening on ", addr)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello World!")
	})

	iris.Listen(addr)
}
