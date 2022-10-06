package entity

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey;type:serial"`
	Items     []CartDetail
	Status    string `gorm:"default:unpaid"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartRequest struct {
	Items []CartDetailRequest `json:"items"`
}

type CartResponse struct {
	ID        uint                 `json:"id"`
	Items     []CartDetailResponse `json:"items"`
	Status    string               `json:"status"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}
