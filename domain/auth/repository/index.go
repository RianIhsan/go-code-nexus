package repository

import (
	"github.com/RianIhsan/go-code-nexus/domain/auth"
	"github.com/RianIhsan/go-code-nexus/entities"
	"gorm.io/gorm"
	"time"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.IRepoAuth {
	return &AuthRepository{
		db,
	}
}

func (r *AuthRepository) InsertToken(token *entities.TokenEntity) (*entities.TokenEntity, error) {
	if err := r.db.Create(&token).Error; err != nil {
		return nil, err
	}
	return token, nil
}

func (r *AuthRepository) SignUp(newUser *entities.UserEntity) (*entities.UserEntity, error) {
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *AuthRepository) FindValidToken(userID int, token string) (*entities.TokenEntity, error) {
	var isValidToken entities.TokenEntity
	if err := r.db.Where("user_id = ? AND token = ? AND expired_token > ?", userID, token, time.Now().Unix()).Find(&isValidToken).Error; err != nil {
		return &isValidToken, err
	}
	return &isValidToken, nil
}

func (r *AuthRepository) UpdateUser(user *entities.UserEntity) (*entities.UserEntity, error) {
	err := r.db.Model(&user).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) DeleteToken(token *entities.TokenEntity) error {
	if err := r.db.Delete(&token).Error; err != nil {
		return err
	}
	return nil
}
