package main

import (
	"fmt"
	"github.com/RianIhsan/go-code-nexus/config"
	_ "github.com/RianIhsan/go-code-nexus/docs"
	hAuth "github.com/RianIhsan/go-code-nexus/domain/auth/handler"
	rAuth "github.com/RianIhsan/go-code-nexus/domain/auth/repository"
	sAuth "github.com/RianIhsan/go-code-nexus/domain/auth/service"
	hUser "github.com/RianIhsan/go-code-nexus/domain/users/handler"
	rUser "github.com/RianIhsan/go-code-nexus/domain/users/repository"
	sUser "github.com/RianIhsan/go-code-nexus/domain/users/service"
	"github.com/RianIhsan/go-code-nexus/middleware"
	"github.com/RianIhsan/go-code-nexus/routes"
	"github.com/RianIhsan/go-code-nexus/utils/database"
	"github.com/RianIhsan/go-code-nexus/utils/hashing"
	jwt2 "github.com/RianIhsan/go-code-nexus/utils/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/fiber-swagger"
)

// @Title CodeNexus API
// @version 1.0
// @description Happy integration
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app := fiber.New(fiber.Config{
		AppName:       "Welcome to code nexus",
		CaseSensitive: false,
	})
	var bootConfig = config.BootConfig()

	db := database.BootDatabase(*bootConfig)
	database.MigrateTable(db)
	hash := hashing.NewHash()
	jwt := jwt2.NewJWT(bootConfig.Token)

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo)
	userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, userService, hash, jwt)
	authHandler := hAuth.NewAuthHandler(authService)

	app.Use(middleware.Logging())

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello Code Nexus 🚀",
		})
	})

	routes.BootRouteAuth(app, authHandler)
	routes.BootRouteUser(app, userHandler, jwt, userService)

	addr := fmt.Sprintf(":%d", bootConfig.AppPort)
	app.Listen(addr)
}
