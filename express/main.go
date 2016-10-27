package main

import (
	"fmt"
	"time"

	"github.com/icebob/go-http-framework-benchmark/util"
	"github.com/icebob/goexpress"
	//"github.com/icebob/goexpress/middlewares"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func registerRoutes(app *goexpress.Express) {
	for _, r := range util.DynamicRoutes {
		switch r.Method {
		case "GET":
			app.Get(r.Path, routeHandler)
		case "POST":
			app.Post(r.Path, routeHandler)
		case "PUT":
			app.Put(r.Path, routeHandler)
		case "PATCH":
			app.Patch(r.Path, routeHandler)
		case "DELETE":
			app.Delete(r.Path, routeHandler)
		default:
			panic("method not supported")
		}
	}
}

func routeHandler(req *request.Request, res *response.Response, next func()) {
	res.Send(fmt.Sprintf("%v", req.Params))
}

func main() {
	//runtime.GOMAXPROCS(8)

	var app = goexpress.New()

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

	app.Get("/slow", func(req *request.Request, res *response.Response, next func()) {
		time.Sleep(10 * time.Millisecond)
		res.Send("Done")
	})

	registerRoutes(app)

	app.Post("/action/:type", func(req *request.Request, res *response.Response, next func()) {
		actionType := req.Params["type"]
		var json util.JsonReq
		if req.JSON.Decode(&json) == nil {
			if actionType == "add" {
				sum := util.Add(json.Num1, json.Num2)
				res.JSON(util.JsonRes{
					Action: actionType,
					Num1:   json.Num1,
					Num2:   json.Num2,
					Result: sum,
				})
			} else {
				res.Error(400, "Not supported action: "+actionType)
			}
		} else {
			res.Error(400, "")
		}
	})

	bindAddress, bindPort := "127.0.0.1", 8002
	fmt.Printf("Listening on %s:%d...\n", bindAddress, bindPort)
	err := app.Listen(bindPort, bindAddress)
	if err != nil {
		panic(err)
	}
}
