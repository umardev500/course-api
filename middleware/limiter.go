package middleware

import (
	"course-api/domain/model"
	"course-api/storage"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:          3,
		Expiration:   1 * time.Minute,
		KeyGenerator: keyGenerator,
		Storage:      storage.NewRedisStorage(),
	})
}

// keyGenerator make custom key for limiter
func keyGenerator(c *fiber.Ctx) string {
	var creds *model.LoginRequest = new(model.LoginRequest)
	if err := c.BodyParser(creds); err != nil {
		return c.IP()
	}

	return creds.User
}
