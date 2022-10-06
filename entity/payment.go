package entity

import (
	"mime/multipart"
	"time"
)

type Payment struct {
	ID        uint `gorm:"primaryKey;type:serial"`
	CartID    uint `gorm:"not null;unique"`
	Cart      Cart
	Photo     string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentRequest struct {
	CartID uint                  `form:"cart_id"`
	Photo  *multipart.FileHeader `form:"photo"`
}

type PaymentResponse struct {
	ID        uint         `json:"id"`
	Cart      CartResponse `json:"cart"`
	Photo     string       `json:"photo"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
