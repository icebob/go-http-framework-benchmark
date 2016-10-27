package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/icebob/go-http-framework-benchmark/util"
	"github.com/valyala/fasthttp"
)

func registerRoutes(e *fasthttprouter.Router) {
	for _, r := range util.DynamicRoutes {
		switch r.Method {
		case "GET":
			e.GET(r.Path, routeHandler)
		case "POST":
			e.POST(r.Path, routeHandler)
		case "PUT":
			e.PUT(r.Path, routeHandler)
		case "PATCH":
			e.PATCH(r.Path, routeHandler)
		case "DELETE":
			e.DELETE(r.Path, routeHandler)
		default:
			panic("method not supported")
		}
	}
}

func routeHandler(c *fasthttp.RequestCtx) {
	fmt.Fprintf(c, "User: %v", c.UserValue("user"))
}

func main() {
	addr := "127.0.0.1:8004"
	fmt.Println("Listening on ", addr)

	router := fasthttprouter.New()
	router.GET("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, "Hello World!")
	})

	router.GET("/slow", func(ctx *fasthttp.RequestCtx) {
		time.Sleep(10 * time.Millisecond)
		fmt.Fprint(ctx, "Done")
	})

	registerRoutes(router)

	router.POST("/action/:type", func(c *fasthttp.RequestCtx) {
		actionType := c.UserValue("type").(string)
		var reqJson util.JsonReq
		if json.Unmarshal(c.PostBody(), &reqJson) == nil {
			if actionType == "add" {
				resValue := util.Add(reqJson.Num1, reqJson.Num2)
				resObj := util.JsonRes{
					Action: actionType,
					Num1:   reqJson.Num1,
					Num2:   reqJson.Num2,
					Result: resValue,
				}
				c.SetContentType("application/json")
				resJson, err := json.Marshal(resObj)
				if err != nil {
					c.Error("", http.StatusInternalServerError)
				} else {
					c.Write(resJson)
				}

			} else {
				c.Error("Not supported action: "+actionType, http.StatusBadRequest)
			}

		} else {
			c.Error("", http.StatusInternalServerError)
		}
	})

	log.Fatal(fasthttp.ListenAndServe(addr, router.Handler))
}
