package controller

import (
	"iris-postgres/model"
	"iris-postgres/repo"

	"github.com/kataras/iris/v12"
)

/*
Lấy danh sách user
*/
func GetAllUser(ctx iris.Context) {
	users, err := repo.GetAllUser()
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
	}

	ctx.JSON(users)
}

/*
Chi tiết thông tin user
*/
func GetUserById(ctx iris.Context) {
	id := ctx.Params().Get("id")
	user, err := repo.GetUserById(id)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(user)
}

/*
Tạo user
*/
func CreateUser(ctx iris.Context) {
	req := new(model.CreateUser)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	user, err := repo.CreateUser(req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(user)
}

/*
Xóa user
*/
func DeleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")

	err := repo.DeleteUser(id)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON("Xóa user thành công")
}

/*
Cập nhật thông tin user
*/
func UpdateUser(ctx iris.Context) {
	id := ctx.Params().Get("id")

	req := new(model.CreateUser)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	user, err := repo.UpdateUser(id, req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(user)
}
