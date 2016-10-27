package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-martini/martini"
	"github.com/icebob/go-http-framework-benchmark/util"
	"github.com/kataras/iris"
)

func registerRoutes(e martini.Router) {
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

func routeHandler(c *iris.Context) {
	c.Write("%v", c.Params)
}

func main() {
	addr := "127.0.0.1:8008"
	fmt.Println("Listening on ", addr)

	m := martini.New()
	router := martini.NewRouter()
	m.Action(router.Handle)
	martini.Env = martini.Prod

	router.Get("/", func(params martini.Params) string {
		return "Hello World!"
	})

	router.Get("/slow", func(params martini.Params) string {
		time.Sleep(10 * time.Millisecond)
		return "Done"
	})

	registerRoutes(router)

	router.Post("/action/:type", func(params martini.Params, w http.ResponseWriter, req *http.Request) {
		actionType := params["type"]
		var reqJson util.JsonReq
		body, _ := ioutil.ReadAll(req.Body)
		if json.Unmarshal(body, &reqJson) == nil {
			if actionType == "add" {
				resValue := util.Add(reqJson.Num1, reqJson.Num2)
				resObj := util.JsonRes{
					Action: actionType,
					Num1:   reqJson.Num1,
					Num2:   reqJson.Num2,
					Result: resValue,
				}
				w.Header().Set("Content-Type", "application/json")
				resJson, err := json.Marshal(resObj)
				if err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				} else {
					w.Write(resJson)
				}

			} else {
				http.Error(w, "Not supported action: "+actionType, http.StatusBadRequest)
			}

		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(addr, m))
}
