package controller

import (
	"course-api/domain/model"

	"github.com/gofiber/fiber/v2"
)

type Auth struct{}

func (a *Auth) Login(c *fiber.Ctx) error {
	var creds *model.LoginRequest = new(model.LoginRequest)
	if err := c.BodyParser(creds); err != nil {
		return failed(c, fiber.StatusUnprocessableEntity, err.Error())
	}

	t := "dummy tokens"
	return ok(c, fiber.StatusOK, "Login successfuly", model.LoginResponse{Token: t})
}

func (a *Auth) Register(c *fiber.Ctx) error {
	return nil
}

func (a *Auth) Logout(c *fiber.Ctx) error {
	return nil
}
