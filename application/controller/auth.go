package controller

import "github.com/gofiber/fiber/v2"

type Auth struct{}

func (a *Auth) Login(c *fiber.Ctx) error {
	return nil
}

func (a *Auth) Register(c *fiber.Ctx) error {
	return nil
}

func (a *Auth) Logout(c *fiber.Ctx) error {
	return nil
}
