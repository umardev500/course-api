package middleware

import (
	"course-api/storage"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        3,
		Expiration: 1 * time.Minute,
		Storage:    storage.NewRedisStorage(),
	})
}
