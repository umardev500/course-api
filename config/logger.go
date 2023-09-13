package config

import "github.com/gofiber/fiber/v2/middleware/logger"

var LoggerConf = logger.Config{
	Format: "[${time}] ${status} - ${method} ${path}\n",
}
