package users

import (
	"github.com/RianIhsan/go-code-nexus/domain/users/dto"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/gofiber/fiber/v2"
)

type IRepoUser interface {
	FindEmail(email string) (*entities.UserEntity, error)
	FindId(id int) (*entities.UserEntity, error)
	InsertUserDetail(user *entities.UserDetailEntity) (*entities.UserDetailEntity, error)
	UpdateUserAvatar(userID int, avatarPath string) error
}

type IServiceUser interface {
	GetEmail(email string) (*entities.UserEntity, error)
	GetId(id int) (*entities.UserEntity, error)
	CreateUserDetail(userID int, request *dto.TCreateUserDetailRequest) (*entities.UserDetailEntity, error)
	UpdateAvatar(userID int, request *dto.TUpdateAvatarRequest) (*entities.UserEntity, error)
}

type IHandlerUsers interface {
	GetCurrentUser(c *fiber.Ctx) error
	CreateUserDetail(c *fiber.Ctx) error
	UpdateAvatar(c *fiber.Ctx) error
}
