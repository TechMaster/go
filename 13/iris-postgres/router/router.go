package router

import (
	"iris-postgres/controller"

	"github.com/kataras/iris/v12"
)

func InitRouter(app *iris.Application) {
	// User router
	app.Get("/users", controller.GetAllUser)
	app.Post("/users", controller.CreateUser)
	app.Get("/users/:id", controller.GetUserById)
	app.Put("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteUser)

	// Post router
	app.Get("/users/:id/posts", controller.GetPostsOfUser)
	app.Post("/users/:id/posts", controller.CreatePost)
	app.Get("/users/:id/posts/:postId", controller.GetPostDetail)
	app.Put("/users/:id/posts/:postId", controller.UpdatePost)
	app.Delete("/users/:id/posts/:postId", controller.DeletePost)
}
