package users

import "github.com/RianIhsan/go-code-nexus/entities"

type IRepoUser interface {
	FindEmail(email string) (*entities.UserEntity, error)
}

type IServiceUser interface {
	GetEmail(email string) (*entities.UserEntity, error)
}
