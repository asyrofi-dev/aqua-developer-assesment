package entity

import "time"

type CartDetail struct {
	ID        uint `gorm:"primaryKey;type:serial"`
	CartID    uint
	ProductID uint
	Product   Product
	Amount    uint `gorm:"not null;type:integer;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartDetailRequest struct {
	ProductID uint `json:"product_id" example:"1"`
	Amount    uint `json:"amount" example:"10"`
}

type CartDetailResponse struct {
	Product   ProductResponse `json:"item"`
	Amount    uint            `json:"amount" example:"10"`
	CreatedAt time.Time       `json:"created_at" example:"2022-10-07T06:49:42.629210158+07:00"`
	UpdatedAt time.Time       `json:"updated_at" example:"2022-10-07T06:49:42.629210158+07:00"`
}
