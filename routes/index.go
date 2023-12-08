package routes

import (
	"github.com/RianIhsan/go-code-nexus/domain/auth"
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/middleware"
	"github.com/RianIhsan/go-code-nexus/utils/jwt"
	"github.com/gofiber/fiber/v2"
)

func BootRouteAuth(app *fiber.App, handler auth.IHandlerAuth) {
	authGroup := app.Group("api/v1/auth")
	authGroup.Post("/signup", handler.SignUp)
	authGroup.Post("/signin", handler.SignIn)
	authGroup.Post("/verify", handler.Verification)
}

func BootRouteUser(app *fiber.App, handler users.IHandlerUsers, jwtService jwt.IJwt, userService users.IServiceUser) {
	userGroup := app.Group("api/v1/user")

	userGroup.Get("/me", middleware.Protected(jwtService, userService), handler.GetCurrentUser)
}
