package main

import (
	"fmt"
	"github.com/RianIhsan/go-code-nexus/config"
	"github.com/RianIhsan/go-code-nexus/middleware"
	"github.com/RianIhsan/go-code-nexus/utils/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to code nexus",
		CaseSensitive: false,
	})
	var bootConfig = config.BootConfig()

	database.BootDatabase(*bootConfig)

	app.Use(middleware.Logging())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello Code Nexus 🚀",
		})
	})

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	app.Listen(addr)
}
