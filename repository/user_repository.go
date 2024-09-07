package repository

import (
	"blogging-platform/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user model.User) (model.User, error)
	FindByUsername(username string) (model.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB}
}

func (ur *userRepository) Register(user model.User) (model.User, error) {
	err := ur.DB.Create(&user).Error
	return user, err
}

func (ur *userRepository) FindByUsername(username string) (model.User, error) {
	var user model.User

	err := ur.DB.Where("username = ?", username).First(&user).Error
	return user, err
}
