package controller

import (
	"github.com/kataras/iris/v12"
)

func ResponseErr(ctx iris.Context, statusCode int, message string) {
	ctx.JSON(iris.Map{
		"status":  statusCode,
		"message": message,
	})
}
