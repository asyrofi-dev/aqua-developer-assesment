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
	ProductID uint `json:"product_id"`
	Amount    uint `json:"amount"`
}

type CartDetailResponse struct {
	Product   ProductResponse `json:"item"`
	Amount    uint            `json:"amount"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
