package main

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

func main() {
	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello, Geektutu")
	//})
	r.POST("/gitlabscanner", func(context *gin.Context) {
		println(reflect.TypeOf(context))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
