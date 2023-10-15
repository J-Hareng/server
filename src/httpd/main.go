package main

import (
	"fmt"
	"net/http"
	"server/src/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Printf("Request \n")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", handler.TestGet())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
