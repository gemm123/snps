package model

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id        uuid.UUID `gorm:"column:id"`
	IdUser    uuid.UUID `gorm:"column:id_user"`
	Total     int       `gorm:"column:total"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type OrderProduct struct {
	Id        uuid.UUID `gorm:"column:id"`
	IdOrder   uuid.UUID `gorm:"column:id_order"`
	IdProduct uuid.UUID `gorm:"column:id_product"`
	Quantity  int       `gorm:"column:quantity"`
	Total     int       `gorm:"column:total"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type OrderProductRequest struct {
	IdProduct uuid.UUID `json:"id_product" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Total     int       `json:"total" binding:"required"`
}

type OrderResponse struct {
	IdOrder      uuid.UUID              `json:"id_order"`
	TotalPrice   int                    `json:"total_price"`
	Status       string                 `json:"status"`
	OrderProduct []OrderProductResponse `json:"order_product"`
}

type OrderProductResponse struct {
	IdProduct uuid.UUID `json:"id_product"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	Total     int       `json:"total"`
}
