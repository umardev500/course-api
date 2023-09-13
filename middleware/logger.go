package middleware

import (
	"course-api/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewLogger() fiber.Handler {
	return logger.New(config.LoggerConf)
}
