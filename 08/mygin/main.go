package main

import (
	"mygin/repo"
	"mygin/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/people", func(c *gin.Context) {
		c.JSON(200, repo.ListPeople())
	})

	router.InitRouter(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
