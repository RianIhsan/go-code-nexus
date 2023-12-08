package service

import (
	"errors"
	"github.com/RianIhsan/go-code-nexus/domain/auth"
	"github.com/RianIhsan/go-code-nexus/domain/auth/dto"
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/utils/email"
	"github.com/RianIhsan/go-code-nexus/utils/hashing"
	"github.com/RianIhsan/go-code-nexus/utils/jwt"
	"github.com/RianIhsan/go-code-nexus/utils/token"
	"time"
)

type AuthService struct {
	repo    auth.IRepoAuth
	SUser   users.IServiceUser
	hashing hashing.HashInterface
	jwt     jwt.IJwt
}

func NewAuthService(repo auth.IRepoAuth, SUser users.IServiceUser, hashing hashing.HashInterface, jwt jwt.IJwt) auth.IServiceAuth {
	return &AuthService{
		repo, SUser, hashing, jwt,
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
		Avatar:   "https://res.cloudinary.com/dyominih0/image/upload/v1702051015/my-sample-avatar/voaa01wefhnziwzqwn1m.webp",
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

func (s *AuthService) Verification(payload *dto.TVerificationRequest) error {
	user, err := s.SUser.GetEmail(payload.Email)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("email not found")
	}
	isValidOTP, err := s.repo.FindValidToken(int(user.ID), payload.Token)
	if err != nil {
		return err
	}

	if isValidOTP.ID == 0 {
		return errors.New("invalid token")
	}

	user.IsVerified = true

	_, errUpdate := s.repo.UpdateUser(user)
	if errUpdate != nil {
		return errors.New("failed verification email")
	}

	errDeleteOTP := s.repo.DeleteToken(isValidOTP)
	if errDeleteOTP != nil {
		return errors.New("failed delete token")
	}

	return nil
}

func (s *AuthService) SignIn(payload *dto.TLoginRequest) (*entities.UserEntity, string, error) {
	user, err := s.SUser.GetEmail(payload.Email)
	if err != nil {
		return nil, "", errors.New("user not found")
	}
	if !user.IsVerified {
		return nil, "", errors.New("your account has not been verified")
	}
	isValidPassword, err := s.hashing.ComparePassword(user.Password, payload.Password)
	if err != nil || !isValidPassword {
		return nil, "", errors.New("incorrect password")
	}

	accessSecret, err := s.jwt.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}
	return user, accessSecret, nil
}
