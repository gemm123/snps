package model

import (
	"github.com/google/uuid"
	"time"
)

type Cart struct {
	Id        uuid.UUID `gorm:"column:id"`
	IdUser    uuid.UUID `gorm:"column:id_user"`
	Total     int       `gorm:"column:total"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CartProduct struct {
	Id        uuid.UUID `gorm:"column:id"`
	IdCart    uuid.UUID `gorm:"column:id_cart"`
	IdProduct uuid.UUID `gorm:"column:id_product"`
	Quantity  int       `gorm:"column:quantity"`
	Total     int       `gorm:"column:total"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type CartRequest struct {
	IdProduct uuid.UUID `json:"id_product" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Total     int       `json:"total" binding:"required"`
}
