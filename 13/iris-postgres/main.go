package main

import (
	"iris-postgres/repo"
	"iris-postgres/router"

	"github.com/kataras/iris/v12"
)

func main() {
	// Connect Database
	db := repo.ConnectDatabase()
	defer db.Close()

	// Init Iris App
	app := iris.New()

	// Register router
	router.InitRouter(app)

	// Iris App listen port
	app.Listen(":3001")
}
