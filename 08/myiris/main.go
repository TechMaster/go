package main

import (
	"myiris/repo"
	"myiris/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Get("/hello", func(ctx iris.Context) {
		_, _ = ctx.WriteString("Hello World")
	})

	app.Get("/ping", func(ctx iris.Context) {
		_, _ = ctx.JSON(iris.Map{"message": "pong"})
	})

	app.Get("/people", func(ctx iris.Context) {
		_, _ = ctx.JSON(repo.ListPeople())
	})

	router.InitRouter(app)

	err := app.Listen(":8080", iris.WithOptimizations)
	if err != nil {
		panic(err)
	}
}
