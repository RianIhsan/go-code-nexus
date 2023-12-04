package auth

import (
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/module/auth/dto"
	"github.com/gofiber/fiber/v2"
)

type IRepoAuth interface {
	SignUp(newUser *entities.UserEntity) (*entities.UserEntity, error)
	InsertToken(token *entities.TokenEntity) (*entities.TokenEntity, error)
}

type IServiceAuth interface {
	SignUp(payload *dto.TRegisterRequest) (*entities.UserEntity, error)
}

type IHandlerAuth interface {
	SignUp(c *fiber.Ctx) error
}
