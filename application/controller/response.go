package controller

import (
	"course-api/domain/model"

	"github.com/gofiber/fiber/v2"
)

func ok(c *fiber.Ctx, status int, message string, data interface{}) error {
	var payload model.OK = model.OK{
		Status:  status,
		Success: true,
		Message: message,
		Data:    data,
	}

	return c.Status(status).JSON(payload)
}

func failed(c *fiber.Ctx, status int, message string) error {
	var payload model.Err = model.Err{
		Status:  status,
		Success: false,
		Message: message,
	}

	return c.Status(status).JSON(payload)
}
