package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/icebob/go-http-framework-benchmark/util"
	"github.com/julienschmidt/httprouter"
)

func registerRoutes(e *httprouter.Router) {
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

func routeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "%v", ps)
}

func main() {
	addr := "127.0.0.1:8006"
	fmt.Println("Listening on ", addr)

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprint(w, "Hello World!")
	})

	router.GET("/slow", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		time.Sleep(10 * time.Millisecond)
		fmt.Fprint(w, "Done")
	})

	registerRoutes(router)

	router.POST("/action/:type", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		actionType := ps.ByName("type")
		var reqJson util.JsonReq
		body, _ := ioutil.ReadAll(r.Body)
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

	log.Fatal(http.ListenAndServe(addr, router))
}
