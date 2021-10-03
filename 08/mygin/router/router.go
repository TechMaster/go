package router

import (
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func InitRouter(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("/login", default_handler)
		v1.GET("/submit", default_handler)
		v1.GET("/read", default_handler)
	}

	v2 := app.Group("/v2")
	{
		for i := 0; i < 40; i++ {
			path, _ := gonanoid.Generate("abcdefgijklmnopqrstuvwxyz0123456789", 8)
			v2.GET("/"+path, default_handler)
		}
		v2.GET("/ox13", default_handler)
	}

}

func default_handler(c *gin.Context) {
	c.String(200, c.FullPath())
}
