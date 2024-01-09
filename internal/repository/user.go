package repository

import (
	"gorm.io/gorm"
	"synapsis/internal/model"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	CheckEmail(email string) (string, error)
	InsertUser(user model.User) error
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{
		DB: DB,
	}
}

func (r *userRepository) CheckEmail(email string) (string, error) {
	var emailDB string
	err := r.DB.Table("users").Select("email").Where("email = ?", email).Find(&emailDB).Error
	return emailDB, err
}

func (r *userRepository) InsertUser(user model.User) error {
	err := r.DB.Table("users").Create(&user).Error
	return err
}
