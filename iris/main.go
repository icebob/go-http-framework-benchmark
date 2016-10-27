package main

import (
	"fmt"
	"time"

	"github.com/icebob/go-http-framework-benchmark/util"
	"github.com/kataras/iris"
)

func registerRoutes() {
	for _, r := range util.DynamicRoutes {
		switch r.Method {
		case "GET":
			iris.Get(r.Path, routeHandler)
		case "POST":
			iris.Post(r.Path, routeHandler)
		case "PUT":
			iris.Put(r.Path, routeHandler)
		case "PATCH":
			iris.Patch(r.Path, routeHandler)
		case "DELETE":
			iris.Delete(r.Path, routeHandler)
		default:
			panic("method not supported")
		}
	}
}

func routeHandler(c *iris.Context) {
	c.Write("%v", c.Params)
}

func main() {
	addr := "127.0.0.1:8007"
	fmt.Println("Listening on ", addr)

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello World!")
	})

	iris.Get("/slow", func(ctx *iris.Context) {
		time.Sleep(10 * time.Millisecond)
		ctx.Write("Done")
	})

	registerRoutes()

	iris.Post("/action/:type", func(c *iris.Context) {
		actionType := c.Param("type")
		var json util.JsonReq
		if c.ReadJSON(&json) == nil {
			if actionType == "add" {
				res := util.Add(json.Num1, json.Num2)
				c.JSON(200, util.JsonRes{
					Action: actionType,
					Num1:   json.Num1,
					Num2:   json.Num2,
					Result: res,
				})
			} else {
				c.EmitError(iris.StatusBadRequest) //, errors.New("Not supported action: "+actionType))
			}
		} else {
			c.EmitError(iris.StatusBadRequest)
		}
	})

	iris.Listen(addr)
}
