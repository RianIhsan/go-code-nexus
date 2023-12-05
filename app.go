package main

import (
	"fmt"
	"github.com/RianIhsan/go-code-nexus/config"
	_ "github.com/RianIhsan/go-code-nexus/docs"
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
	"github.com/swaggo/fiber-swagger"
)

// @Title CodeNexus API
// @version 1.0
// @description Happy integration
// @host localhost:3000
// @BasePath /
// @schemes http
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

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello Code Nexus 🚀",
		})
	})

	routes.BootRouteAuth(app, authHandler)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	app.Listen(addr)
}
