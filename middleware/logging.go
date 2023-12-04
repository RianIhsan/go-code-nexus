package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logging() fiber.Handler {
	return logger.New(logger.Config{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
	})
}
