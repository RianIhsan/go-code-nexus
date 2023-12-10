package repository

import (
	"github.com/RianIhsan/go-code-nexus/domain/users"
	"github.com/RianIhsan/go-code-nexus/entities"
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

func (r *UserRepository) FindId(id int) (*entities.UserEntity, error) {
	var user *entities.UserEntity
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) InsertUserDetail(user *entities.UserDetailEntity) (*entities.UserDetailEntity, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUserAvatar(userID int, avatarPath string) error {
	user := &entities.UserEntity{ID: userID}

	// Memperbarui hanya field Avatar
	result := r.db.Model(user).Updates(map[string]interface{}{"avatar": avatarPath})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
