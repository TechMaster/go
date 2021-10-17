package controller

import (
	"fiber-postgres/model"
	"fiber-postgres/repo"

	"github.com/gofiber/fiber/v2"
)

/*
Lấy danh sách post của user
*/
func GetPostsOfUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	posts, err := repo.GetPostsOfUser(userId)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(posts)
}

/*
Lấy thông tin của post
*/
func GetPostDetail(c *fiber.Ctx) error {
	userId := c.Params("id")
	postId := c.Params("postId")

	post, err := repo.GetPostDetail(userId, postId)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(post)
}

/*
Tạo post
*/
func CreatePost(c *fiber.Ctx) error {
	userId := c.Params("id")

	req := new(model.CreatePost)
	if err := c.BodyParser(req); err != nil {
		return ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}

	post, err := repo.CreatePost(userId, req)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(post)
}

/*
Xóa post
*/
func DeletePost(c *fiber.Ctx) error {
	userId := c.Params("id")
	postId := c.Params("postId")

	err := repo.DeletePost(userId, postId)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON("Xóa post thành công")
}

/*
Update post
*/
func UpdatePost(c *fiber.Ctx) error {
	userId := c.Params("id")
	postId := c.Params("postId")

	req := new(model.CreatePost)
	if err := c.BodyParser(req); err != nil {
		return ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}

	post, err := repo.UpdatePost(userId, postId, req)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(post)
}
