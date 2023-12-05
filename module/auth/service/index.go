package service

import (
	"errors"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/module/auth"
	"github.com/RianIhsan/go-code-nexus/module/auth/dto"
	"github.com/RianIhsan/go-code-nexus/module/users"
	"github.com/RianIhsan/go-code-nexus/utils/email"
	"github.com/RianIhsan/go-code-nexus/utils/hashing"
	"github.com/RianIhsan/go-code-nexus/utils/token"
	"time"
)

type AuthService struct {
	repo    auth.IRepoAuth
	SUser   users.IServiceUser
	hashing hashing.HashInterface
}

func NewAuthService(repo auth.IRepoAuth, SUser users.IServiceUser, hashing hashing.HashInterface) auth.IServiceAuth {
	return &AuthService{
		repo, SUser, hashing,
	}
}

func (s *AuthService) SignUp(payload *dto.TRegisterRequest) (*entities.UserEntity, error) {
	isExistUser, _ := s.SUser.GetEmail(payload.Email)
	if isExistUser != nil {
		return nil, errors.New("email already exists")
	}

	hashPassword, err := s.hashing.GenerateHash(payload.Password)
	if err != nil {
		return nil, err
	}

	newUser := &entities.UserEntity{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashPassword,
		Role:     "user",
	}

	user, err := s.repo.SignUp(newUser)
	if err != nil {
		return nil, errors.New("failed create account")
	}

	newToken := token.GenerateRandomOTP(6)
	saveToken := &entities.TokenEntity{
		UserID:       user.ID,
		Token:        newToken,
		ExpiredToken: time.Now().Add(2 * time.Minute).Unix(),
	}

	_, errToken := s.repo.InsertToken(saveToken)
	if errToken != nil {
		return nil, errors.New("failed save token email")
	}

	err = email.EmailService(user.Name, user.Email, newToken)
	if err != nil {
		return nil, errors.New("failed send token to email")
	}
	return user, nil
}
