package main

import (
	"fmt"

	"github.com/icebob/goexpress"
	//"github.com/icebob/goexpress/middlewares"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func main() {
	//runtime.GOMAXPROCS(8)

	var app = goexpress.Express()

	//app.Use(middlewares.Log(middlewares.LOGTYPE_DEV))
	/*
		app.Get("/test", func(req *request.Request, res *response.Response, next func()) {
			res.Send("Hello Test")
		})

		app.Get("/chunked", func(req *request.Request, res *response.Response, next func()) {
			res.WriteChunk("Hello World\n")
			//time.Sleep(5 * time.Second)
			res.WriteChunk("Hello World2\n")
		})*/

	app.Get("/", func(req *request.Request, res *response.Response, next func()) {
		res.Send("Hello World!")
	})

	bindAddress, bindPort := "127.0.0.1", 8002
	fmt.Printf("Listening on %s:%d...\n", bindAddress, bindPort)
	err := app.Listen(bindPort, bindAddress)
	if err != nil {
		panic(err)
	}
}
