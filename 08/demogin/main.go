package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//r := gin.Default()
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":8080")
}
