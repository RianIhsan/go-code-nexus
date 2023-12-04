package main

import (
	"fmt"
	"github.com/RianIhsan/go-code-nexus/config"
	"github.com/RianIhsan/go-code-nexus/middleware"
	"github.com/RianIhsan/go-code-nexus/module/auth/handler"
	rAuth "github.com/RianIhsan/go-code-nexus/module/auth/repository"
	sAuth "github.com/RianIhsan/go-code-nexus/module/auth/service"
	rUser "github.com/RianIhsan/go-code-nexus/module/users/repository"
	sUser "github.com/RianIhsan/go-code-nexus/module/users/service"
	"github.com/RianIhsan/go-code-nexus/routes"
	"github.com/RianIhsan/go-code-nexus/utils/database"
	"github.com/RianIhsan/go-code-nexus/utils/hashing"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to code nexus",
		CaseSensitive: false,
	})
	var bootConfig = config.BootConfig()

	db := database.BootDatabase(*bootConfig)
	database.MigrateTable(db)
	hash := hashing.NewHash()

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, userService, hash)
	authHandler := handler.NewAuthHandler(authService)

	app.Use(middleware.Logging())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello Code Nexus 🚀",
		})
	})

	routes.BootRouteAuth(app, authHandler)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	app.Listen(addr)
}
