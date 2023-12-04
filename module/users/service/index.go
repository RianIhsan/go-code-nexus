package service

import (
	"errors"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/module/users"
)

type UserService struct {
	repo users.IRepoUser
}

func NewUserService(repo users.IRepoUser) users.IServiceUser {
	return &UserService{
		repo,
	}
}

func (s *UserService) GetEmail(email string) (*entities.UserEntity, error) {
	result, err := s.repo.FindEmail(email)
	if err != nil {
		return nil, errors.New("Your email has been already")
	}
	return result, nil
}
