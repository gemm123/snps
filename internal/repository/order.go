package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"synapsis/internal/model"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	InsertOrder(order model.Order) error
	InsertOrderProduct(orderProduct model.OrderProduct) error
	GetAllOrderProductByIdOrder(idOrder uuid.UUID) ([]model.OrderProduct, error)
	GetOrdersByIdUser(idUser uuid.UUID) ([]model.Order, error)
	GetOrderById(idOrder uuid.UUID) (model.Order, error)
	UpdateOrder(order model.Order) error
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

func (r *orderRepository) GetAllOrderProductByIdOrder(idOrder uuid.UUID) ([]model.OrderProduct, error) {
	var orderProduct []model.OrderProduct
	err := r.DB.Table("order_products").Where("id_order = ?", idOrder).Find(&orderProduct).Error
	return orderProduct, err
}

func (r *orderRepository) GetOrdersByIdUser(idUser uuid.UUID) ([]model.Order, error) {
	var order []model.Order
	err := r.DB.Table("orders").Where("id_user = ?", idUser).Find(&order).Error
	return order, err
}

func (r *orderRepository) GetOrderById(idOrder uuid.UUID) (model.Order, error) {
	var order model.Order
	err := r.DB.Table("orders").Where("id = ?", idOrder).Find(&order).Error
	return order, err
}

func (r *orderRepository) UpdateOrder(order model.Order) error {
	err := r.DB.Save(&order).Error
	return err
}
