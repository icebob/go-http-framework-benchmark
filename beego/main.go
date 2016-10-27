package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/icebob/go-http-framework-benchmark/util"
)

func registerRoutes(app *beego.ControllerRegister) {
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

func routeHandler(c *context.Context) {
	c.WriteString(fmt.Sprintf("%v", c.Input.Params))
}

func main() {
	addr := "127.0.0.1:8009"
	fmt.Println("Listening on ", addr)

	beego.BConfig.RunMode = beego.PROD
	beego.BeeLogger.Close()
	app := beego.NewControllerRegister()

	app.Get("/", func(c *context.Context) {
		c.WriteString("Hello World!")
	})

	app.Get("/slow", func(c *context.Context) {
		time.Sleep(10 * time.Millisecond)
		c.WriteString("Done")
	})

	registerRoutes(app)

	app.Post("/action/:type", func(c *context.Context) {
		actionType := c.Input.Param(":type")
		var reqJson util.JsonReq
		body := c.Input.CopyBody(1000)
		if json.Unmarshal(body, &reqJson) == nil {
			if actionType == "add" {
				resValue := util.Add(reqJson.Num1, reqJson.Num2)
				resObj := util.JsonRes{
					Action: actionType,
					Num1:   reqJson.Num1,
					Num2:   reqJson.Num2,
					Result: resValue,
				}
				c.Output.ContentType("application/json")
				c.ResponseWriter.Header().Set("Content-Type", "application/json")
				resJson, err := json.Marshal(resObj)
				if err != nil {
					http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				} else {
					c.ResponseWriter.Write(resJson)
				}

			} else {
				http.Error(c.ResponseWriter, "Not supported action: "+actionType, http.StatusBadRequest)
			}

		} else {
			http.Error(c.ResponseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(addr, app))
}
