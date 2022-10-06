package entity

import (
	"mime/multipart"
	"time"
)

type Product struct {
	ID          uint   `gorm:"primaryKey;type:serial"`
	Name        string `gorm:"not null;unique"`
	Photo       string `gorm:"not null"`
	Price       uint   `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRequest struct {
	Name        string                `form:"name"`
	Photo       *multipart.FileHeader `form:"photo"`
	Price       uint                  `form:"price"`
	Description string                `form:"desc"`
}

type ProductResponse struct {
	Name        string    `json:"name"`
	Photo       string    `json:"photo"`
	Price       uint      `json:"price"`
	Description string    `json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
