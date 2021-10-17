package controller

import (
	"iris-postgres/model"
	"iris-postgres/repo"

	"github.com/kataras/iris/v12"
)

/*
Lấy danh sách post của user
*/
func GetPostsOfUser(ctx iris.Context) {
	userId := ctx.Params().Get("id")
	posts, err := repo.GetPostsOfUser(userId)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(posts)
}

/*
Lấy thông tin của post
*/
func GetPostDetail(ctx iris.Context) {
	userId := ctx.Params().Get("id")
	postId := ctx.Params().Get("postId")

	post, err := repo.GetPostDetail(userId, postId)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(post)
}

/*
Tạo post
*/
func CreatePost(ctx iris.Context) {
	userId := ctx.Params().Get("id")

	req := new(model.CreatePost)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	post, err := repo.CreatePost(userId, req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(post)
}

/*
Xóa post
*/
func DeletePost(ctx iris.Context) {
	userId := ctx.Params().Get("id")
	postId := ctx.Params().Get("postId")

	err := repo.DeletePost(userId, postId)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON("Xóa post thành công")
}

/*
Update post
*/
func UpdatePost(ctx iris.Context) {
	userId := ctx.Params().Get("id")
	postId := ctx.Params().Get("postId")

	req := new(model.CreatePost)
	if err := ctx.ReadJSON(req); err != nil {
		ResponseErr(ctx, iris.StatusBadRequest, err.Error())
		return
	}

	post, err := repo.UpdatePost(userId, postId, req)
	if err != nil {
		ResponseErr(ctx, iris.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(post)
}
