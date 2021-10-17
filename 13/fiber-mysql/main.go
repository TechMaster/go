package main

import (
	"fiber-postgres/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Init Fiber App
	app := fiber.New()

	// Register router
	router.InitRouter(app)

	// Fiber App listen port
	app.Listen(":3002")
}
