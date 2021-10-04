package main

import (
	"errors"
	"log"
	"myiris/repo"
	"myiris/router"
	"net/http"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	//app.Use(logger.New(logger.DefaultConfig()))
	//app.UseRouter(recover.New())

	app.Get("/hello", func(ctx iris.Context) {
		_, _ = ctx.WriteString("Hello World")
	})

	app.Get("/ping", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Get("/people", func(ctx iris.Context) {
		_, _ = ctx.JSON(repo.ListPeople())
	})

	app.Post("/book", func(ctx iris.Context) {
		var book repo.Book
		if err := ctx.ReadJSON(&book); err == nil {
			log.Println(book.Title)
			log.Println(book.Author)
			log.Println(book.Pages)
			_, _ = ctx.WriteString("Success")
		} else {
			ctx.StatusCode(http.StatusBadRequest)
			_, _ = ctx.WriteString(err.Error())
		}
	})

	app.Get("/panic", func(ctx iris.Context) {
		panic(errors.New("Make silly error"))
	})

	router.InitRouter(app)

	err := app.Listen(":8080", iris.WithOptimizations)
	if err != nil {
		panic(err)
	}
}
