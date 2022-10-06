package entity

import (
	"mime/multipart"
	"time"
)

type Product struct {
	ID          uint   `gorm:"primaryKey;type:serial"`
	Name        string `gorm:"not null;unique;default:null"`
	Category    string `gorm:"not null;default:null"`
	Photo       string `gorm:"not null"`
	Price       uint   `gorm:"not null;type:integer;default:null"`
	Description string `gorm:"not null;default:null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductRequest struct {
	Name        string                `form:"name"`
	Category    string                `form:"category"`
	Photo       *multipart.FileHeader `form:"photo"`
	Price       uint                  `form:"price"`
	Description string                `form:"desc"`
}

type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Photo       string    `json:"photo"`
	Price       uint      `json:"price"`
	Description string    `json:"desc"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProductFilter struct {
	Category string `query:"category"`
	MinPrice int    `query:"minprice"`
	MaxPrice int    `query:"maxprice"`
}
