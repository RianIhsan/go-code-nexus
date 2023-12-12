package service

import (
	"errors"
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/domain/users/dto"
	"github.com/RianIhsan/go-code-nexus/entities"
	"time"
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

func (s *UserService) CreateUserDetail(userID int, request *dto.TCreateUserDetailRequest) (*entities.UserDetailEntity, error) {
	user, err := s.repo.FindId(userID)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	userDetail := &entities.UserDetailEntity{
		UserID:    userID,
		Address:   request.Address,
		Phone:     request.Phone,
		Job:       request.Job,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userDetail, err = s.repo.InsertUserDetail(userDetail)
	if err != nil {
		return nil, errors.New("failed to create user detail")
	}

	return userDetail, nil
}

func (s *UserService) UpdateAvatar(userID int, request *dto.TUpdateAvatarRequest) (*entities.UserEntity, error) {
	// Mendapatkan data pengguna dari database
	user, err := s.repo.FindId(userID)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	// Membuat objek user baru dengan hanya memperbarui field Avatar
	userUpdateAvatar := &entities.UserEntity{
		Avatar: request.Avatar,
	}

	// Memperbarui avatar pada database
	err = s.repo.UpdateUserAvatar(userID, userUpdateAvatar.Avatar)
	if err != nil {
		return nil, errors.New("failed to update user avatar")
	}

	// Memperbarui objek pengguna dengan data baru
	user.Avatar = userUpdateAvatar.Avatar

	return user, nil
}
