package service

import (
	"errors"
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/entities"
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

func (s *UserService) GetId(id int) (*entities.UserEntity, error) {
	user, err := s.repo.FindId(id)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}
