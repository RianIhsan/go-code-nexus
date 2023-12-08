package users

import (
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/gofiber/fiber/v2"
)

type IRepoUser interface {
	FindEmail(email string) (*entities.UserEntity, error)
	FindId(id int) (*entities.UserEntity, error)
}

type IServiceUser interface {
	GetEmail(email string) (*entities.UserEntity, error)
	GetId(id int) (*entities.UserEntity, error)
}

type IHandlerUsers interface {
	GetCurrentUser(c *fiber.Ctx) error
}
