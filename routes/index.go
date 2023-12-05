package routes

import (
	"github.com/RianIhsan/go-code-nexus/module/auth"
	"github.com/gofiber/fiber/v2"
)

func BootRouteAuth(app *fiber.App, handler auth.IHandlerAuth) {
	authGroup := app.Group("api/v1/auth")
	authGroup.Post("/signup", handler.SignUp)
}
