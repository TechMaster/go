package router

import (
	"github.com/kataras/iris/v12"
	gonanoid "github.com/matoous/go-nanoid"
)

func InitRouter(app *iris.Application) {
	v1 := app.Party("/v1")
	{
		v1.Get("/login", default_handler)
		v1.Get("/submit", default_handler)
		v1.Get("/read", default_handler)
	}

	v2 := app.Party("/v2")
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

func default_handler(c iris.Context) {
	_, _ = c.WriteString(c.Path())
}
