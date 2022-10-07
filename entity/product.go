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
	Name        string                `form:"name" example:"Pakan Lele"`
	Category    string                `form:"category" example:"pakan"`
	Photo       *multipart.FileHeader `form:"photo" swaggertype:"string" example:"pakan_lele.jpeg"`
	Price       uint                  `form:"price" example:"50000"`
	Description string                `form:"desc" example:"pakan lele berkualitas tinggi"`
}

type ProductResponse struct {
	ID          uint      `json:"id" example:"1"`
	Name        string    `json:"name" example:"Pakan Lele"`
	Category    string    `json:"category" example:"pakan"`
	Photo       string    `json:"photo" example:"/api/v1/images/products/product_1.jpeg"`
	Price       uint      `json:"price" example:"50000"`
	Description string    `json:"desc" example:"Pakan lele berkualitas tinggi"`
	CreatedAt   time.Time `json:"created_at" example:"2022-10-07T06:49:42.629210158+07:00"`
	UpdatedAt   time.Time `json:"updated_at" example:"2022-10-07T06:49:42.629210158+07:00"`
}

type ProductFilter struct {
	Category string `query:"category"`
	MinPrice int    `query:"minprice"`
	MaxPrice int    `query:"maxprice"`
}
