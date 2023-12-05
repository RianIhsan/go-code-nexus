package repository

import (
	"github.com/RianIhsan/go-code-nexus/entities"
	"github.com/RianIhsan/go-code-nexus/module/users"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.IRepoUser {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) FindEmail(email string) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Table("users").
		Where("email = ? AND deleted_at IS NULL", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}
