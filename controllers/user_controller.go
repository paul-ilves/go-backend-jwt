package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/paul-ilves/wanaku-api-go/services"
)

func HandleGetAllUsers(c *fiber.Ctx) error {
	response, appErr := services.GetAllUsers()
	if appErr != nil {
		return c.Status(int(appErr.Code)).JSON(appErr.AsMessage())
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func HandleGetUser(c *fiber.Ctx) error {
	i := c.Params("userID")
	var id uint64
	_, err := fmt.Sscan(i, &id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "Expected number in path variable"})
	}
	response, appErr := services.GetUser(id)
	if appErr != nil {
		return c.Status(int(appErr.Code)).JSON(appErr.AsMessage())
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

