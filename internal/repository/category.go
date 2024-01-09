package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

type CategoryRepository interface {
	GetNameCategoryByIdCategory(idCategory uuid.UUID) (string, error)
}

func NewCategoryRepository(DB *gorm.DB) *categoryRepository {
	return &categoryRepository{
		DB: DB,
	}
}

func (r *categoryRepository) GetNameCategoryByIdCategory(idCategory uuid.UUID) (string, error) {
	var categoryName string
	err := r.DB.Table("categories").
		Select("name").
		Where("id = ?", idCategory).
		Find(&categoryName).Error
	return categoryName, err
}
