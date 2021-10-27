package main

import (
	"uselogger/router"

	logger "github.com/TechMaster/coolog"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	logFile := logger.Init() //Cần phải có 2 file error.html và info.html ở /views
	if logFile != nil {
		defer logFile.Close()
	}
	router.RegisterRoute(app)

	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	_ = app.Listen(":8888")
}
