package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/icebob/go-http-framework-benchmark/util"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func registerRoutes(e *echo.Echo) {
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

func routeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func main() {
	addr := "127.0.0.1:8001"
	fmt.Println("Listening on ", addr)

	e := echo.New()

	// e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/slow", func(c echo.Context) error {
		time.Sleep(10 * time.Millisecond)
		return c.String(http.StatusOK, "Done")
	})

	registerRoutes(e)

	e.POST("/action/:type", func(c echo.Context) error {
		actionType := c.Param("type")
		var json util.JsonReq
		if c.Bind(&json) == nil {
			if actionType == "add" {
				res := util.Add(json.Num1, json.Num2)
				return c.JSON(200, util.JsonRes{
					Action: actionType,
					Num1:   json.Num1,
					Num2:   json.Num2,
					Result: res,
				})
			}
			return echo.NewHTTPError(http.StatusBadRequest, "Not supported action: "+actionType)

		}
		return echo.NewHTTPError(http.StatusInternalServerError)
	})

	e.Run(standard.New(addr))
}
