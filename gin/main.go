package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/icebob/go-http-framework-benchmark/util"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func registerRoutes(e *gin.Engine) {
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

func routeHandler(c *gin.Context) {
	c.String(http.StatusOK, "%v", c.Params)
}

func main() {
	addr := "127.0.0.1:8005"
	fmt.Println("Listening on ", addr)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	//r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	r.GET("/slow", func(c *gin.Context) {
		time.Sleep(10 * time.Millisecond)
		c.String(http.StatusOK, "Done")
	})

	registerRoutes(r)

	r.POST("/action/:type", func(c *gin.Context) {
		actionType := c.Param("type")
		var json util.JsonReq
		if c.BindJSON(&json) == nil {
			if actionType == "add" {
				res := util.Add(json.Num1, json.Num2)
				c.JSON(200, util.JsonRes{
					Action: actionType,
					Num1:   json.Num1,
					Num2:   json.Num2,
					Result: res,
				})
			} else {
				c.AbortWithError(400, errors.New("Not supported action: "+actionType))
			}
		} else {
			c.AbortWithStatus(400)
		}
	})

	//r.Static("/app", "./../static/public")
	//r.StaticFS("/app", http.Dir("./../static/public"))
	r.Run(addr)
}
