package repository

import (
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/module/auth"
	"gorm.io/gorm"
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
