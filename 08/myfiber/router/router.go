package router

import (
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid"
)

func InitRouter(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		v1.Get("/login", default_handler)
		v1.Get("/submit", default_handler)
		v1.Get("/read", default_handler)
	}

	v2 := app.Group("/v2")
	{
		for i := 0; i < 40; i++ {
			path, _ := gonanoid.Generate("abcdefgijklmnopqrstuvwxyz0123456789", 8)
			v2.Get("/"+path, default_handler)
		}
		v2.Get("/ox13", default_handler)

		//Only Iris support regex
		//http://localhost:8080/v2/regex/abc and http://localhost:8080/v2/regex/abc123
		v2.Get("/regex/{name:string regexp(^[a-z]+)}", default_handler)
	}

}

func default_handler(c *fiber.Ctx) (err error) {
	_, err = c.WriteString(c.Path())
	return
}
