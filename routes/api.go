package routes

import (
	"course-api/application/controller"
	"course-api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoadAPIRoutes(app *fiber.App) {
	router := app.Group("/api")
	router.Use(logger.New(config.LoggerConf))
	router.Route("/auth", loadAuthRoutes)
}

func loadAuthRoutes(router fiber.Router) {
	handler := &controller.Auth{
		Validate: config.Validate,
	}

	router.Post("/login", handler.Login)
	router.Post("/register", handler.Register)
	router.Post("/logout", handler.Logout)
}
