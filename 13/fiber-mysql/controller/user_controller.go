package controller

import (
	"fiber-postgres/model"
	"fiber-postgres/repo"

	"github.com/gofiber/fiber/v2"
)

/*
Lấy danh sách user
*/
func GetAllUser(c *fiber.Ctx) error {
	users, err := repo.GetAllUser()
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(users)
}

/*
Chi tiết thông tin user
*/
func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := repo.GetUserById(id)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}

/*
Tạo user
*/
func CreateUser(c *fiber.Ctx) error {
	req := new(model.CreateUser)
	if err := c.BodyParser(req); err != nil {
		return ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := repo.CreateUser(req)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}

/*
Xóa user
*/
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := repo.DeleteUser(id)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON("Xóa user thành công")
}

/*
Cập nhật thông tin user
*/
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	req := new(model.CreateUser)
	if err := c.BodyParser(req); err != nil {
		return ResponseErr(c, fiber.StatusBadRequest, err.Error())
	}

	user, err := repo.UpdateUser(id, req)
	if err != nil {
		return ResponseErr(c, fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}
