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
	CartID uint                  `form:"cart_id" example:"1"`
	Photo  *multipart.FileHeader `form:"photo" swaggertype:"string" example:"transfer.jpeg"`
}

type PaymentResponse struct {
	ID        uint         `json:"id" example:"1"`
	Cart      CartResponse `json:"cart"`
	Photo     string       `json:"photo" example:"/api/v1/images/payments/payment_1.jpeg"`
	CreatedAt time.Time    `json:"created_at" example:"2022-10-07T06:49:42.629210158+07:00"`
	UpdatedAt time.Time    `json:"updated_at" example:"2022-10-07T06:49:42.629210158+07:00"`
}
