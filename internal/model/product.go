package model

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID `gorm:"column:id"`
	IdCategory  uuid.UUID `gorm:"column:id_category"`
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	Price       int       `gorm:"column:price"`
	Stok        int       `gorm:"column:stok"`
	Image       string    `gorm:"column:image"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type ProductResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Price       int    `json:"price"`
	Stok        int    `json:"stok"`
	Image       string `json:"image"`
}
