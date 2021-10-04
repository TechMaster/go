package main

import (
	"errors"
	"log"
	"mygin/repo"
	"mygin/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	//r.Use(gin.Logger())

	//r.Use(gin.Recovery())

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

	r.POST("/book", func(c *gin.Context) {
		var book repo.Book
		if err := c.ShouldBind(&book); err == nil {
			log.Println(book.Title)
			log.Println(book.Author)
			log.Println(book.Pages)
		} else {
			log.Println(err.Error())
		}
		c.String(200, "Success")
	})

	//Cố tình tạo lỗi panic
	r.GET("/panic", func(c *gin.Context) {
		//log.Fatalln("Lỗi nghiêm trọng")
		panic(errors.New("Make silly error"))
	})

	router.InitRouter(r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
