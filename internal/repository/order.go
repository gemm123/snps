package repository

import (
	"gorm.io/gorm"
	"synapsis/internal/model"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	InsertOrder(order model.Order) error
	InsertOrderProduct(orderProduct model.OrderProduct) error
}

func NewOrderRepository(DB *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (r *orderRepository) InsertOrder(order model.Order) error {
	err := r.DB.Table("orders").Create(&order).Error
	return err
}

func (r *orderRepository) InsertOrderProduct(orderProduct model.OrderProduct) error {
	err := r.DB.Table("order_products").Create(&orderProduct).Error
	return err
}
