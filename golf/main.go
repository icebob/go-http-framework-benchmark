package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dinever/golf"
	"github.com/icebob/go-http-framework-benchmark/util"
)

func registerRoutes(e *golf.Application) {
	for _, r := range util.DynamicRoutes {
		switch r.Method {
		case "GET":
			e.Get(r.Path, routeHandler)
		case "POST":
			e.Post(r.Path, routeHandler)
		case "PUT":
			e.Put(r.Path, routeHandler)
		case "PATCH":
			e.Patch(r.Path, routeHandler)
		case "DELETE":
			e.Delete(r.Path, routeHandler)
		default:
			panic("method not supported")
		}
	}
}

func routeHandler(ctx *golf.Context) {
	ctx.Send(fmt.Sprintf("%v", ctx.Params))
}

func main() {
	addr := "127.0.0.1:8010"
	fmt.Println("Listening on ", addr)

	app := golf.New()
	app.Get("/", func(ctx *golf.Context) {
		ctx.Send("Hello World!")
	})

	app.Get("/slow", func(ctx *golf.Context) {
		time.Sleep(10 * time.Millisecond)
		ctx.Send("Done!")
	})

	registerRoutes(app)

	app.Post("/action/:type", func(ctx *golf.Context) {
		actionType := ctx.Param("type")
		var reqJson util.JsonReq
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		if json.Unmarshal(body, &reqJson) == nil {
			if actionType == "add" {
				resValue := util.Add(reqJson.Num1, reqJson.Num2)
				resObj := util.JsonRes{
					Action: actionType,
					Num1:   reqJson.Num1,
					Num2:   reqJson.Num2,
					Result: resValue,
				}
				ctx.JSON(resObj)
			} else {
				ctx.Abort(http.StatusBadRequest)
			}

		} else {
			ctx.Abort(http.StatusInternalServerError)
		}
	})

	app.Run(addr)
}
