package router

import (
	"fiber-postgres/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App) {
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
