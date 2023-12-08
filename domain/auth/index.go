package auth

import (
	"github.com/RianIhsan/go-code-nexus/domain/auth/dto"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/gofiber/fiber/v2"
)

type IRepoAuth interface {
	SignUp(newUser *entities.UserEntity) (*entities.UserEntity, error)
	InsertToken(token *entities.TokenEntity) (*entities.TokenEntity, error)
	FindValidToken(userID int, token string) (*entities.TokenEntity, error)
	UpdateUser(user *entities.UserEntity) (*entities.UserEntity, error)
	DeleteToken(token *entities.TokenEntity) error
}

type IServiceAuth interface {
	SignUp(payload *dto.TRegisterRequest) (*entities.UserEntity, error)
	SignIn(payload *dto.TLoginRequest) (*entities.UserEntity, string, error)
	Verification(payload *dto.TVerificationRequest) error
}

type IHandlerAuth interface {
	SignUp(c *fiber.Ctx) error
	Verification(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
}
