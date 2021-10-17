package controller

import "github.com/gofiber/fiber/v2"

func ResponseErr(c *fiber.Ctx, statusCode int, message string) (error){
	return c.JSON(fiber.Map{
		"status":  statusCode,
		"message": message,
	})
}
