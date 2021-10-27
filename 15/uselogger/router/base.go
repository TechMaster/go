package router

import (
	"uselogger/controller"

	"github.com/kataras/iris/v12"
)

func RegisterRoute(app *iris.Application) {
	app.Get("/", controller.Get_a_book)
	app.Get("/books", controller.Get_books)
	app.Get("/info", controller.Process_info)
}
