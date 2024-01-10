package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"synapsis/internal/model"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	GetAllProductByCategoryName(name string) ([]model.Product, error)
	GetProductById(idProduct uuid.UUID) (model.Product, error)
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

func (r *productRepository) GetProductById(idProduct uuid.UUID) (model.Product, error) {
	var product model.Product
	err := r.DB.Table("products").Where("id = ?", idProduct).Find(&product).Error
	return product, err
}
