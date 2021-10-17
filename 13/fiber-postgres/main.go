package main

import (
	"fiber-postgres/repo"
	"fiber-postgres/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect Database
	db := repo.ConnectDatabase()
	defer db.Close()

	// Init Fiber App
	app := fiber.New()

	// Register router
	router.InitRouter(app)

	// Fiber App listen port
	app.Listen(":3000")
}
