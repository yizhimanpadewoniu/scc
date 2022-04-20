package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello, Geektutu")
	//})
	r.POST("/gitlabscanner", events)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func events(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	resp := map[string]string{"hello": "world"}
	c.JSON(http.StatusOK, resp)
	/*post_gwid := c.PostForm("name")
	  fmt.Println(post_gwid)*/

}
