package repository

import (
	"gorm.io/gorm"
	"synapsis/internal/model"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetAllProductByCategoryName(name string) ([]model.Product, error)
}

func NewProductRepository(DB *gorm.DB) *productRepository {
	return &productRepository{
		DB: DB,
	}
}

func (r *productRepository) GetAllProductByCategoryName(name string) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Table("products").
		Joins("left join categories on categories.id = products.id_category").
		Where("categories.name ILIKE ?", name).
		Find(&products).
		Error
	return products, err
}
