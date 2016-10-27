package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/icebob/go-http-framework-benchmark/util"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func main() {
	addr := "127.0.0.1:8005"
	fmt.Println("Listening on ", addr)

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	//r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
		/*c.JSON(200, gin.H{
			"message": "pong",
		})*/
	})

	r.POST("/action", func(c *gin.Context) {
		var json util.JsonReq
		if c.BindJSON(&json) == nil {
			if json.Action == "add" {
				res := util.Add(json.Num1, json.Num2)
				c.JSON(200, util.JsonRes{
					Action: json.Action,
					Num1:   json.Num1,
					Num2:   json.Num2,
					Result: res,
				})
			} else {
				c.AbortWithError(400, errors.New("Not supported action: "+json.Action))
			}
		} else {
			c.AbortWithStatus(400)
		}
	})

	//r.Static("/app", "./../static/public")
	//r.StaticFS("/app", http.Dir("./../static/public"))
	r.Run(addr)
}
