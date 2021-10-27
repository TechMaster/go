package controller

import (
	"uselogger/repo"

	logger "github.com/TechMaster/coolog"
	"github.com/TechMaster/eris"
	"github.com/kataras/iris/v12"
)

func Get_a_book(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"title":  "Gone with the wind",
		"author": "Magaret Mitchel",
	})
}

func Get_books(ctx iris.Context) {
	books, err := repo.Get_books()
	if err != nil {
		logger.Log(ctx, err)
	}
	_, _ = ctx.JSON(books)
}

func Process_info(ctx iris.Context) {
	err := eris.Warning("Email is invalid")
	logger.Log(ctx, err)
}
