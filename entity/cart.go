package entity

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey;type:serial"`
	Items     []CartDetail
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartRequest struct {
	Items []CartDetailRequest `json:"items"`
}

type CartResponse struct {
	ID        uint                 `json:"id" example:"1"`
	Items     []CartDetailResponse `json:"items"`
	CreatedAt time.Time            `json:"created_at" example:"2022-10-07T06:49:42.629210158+07:00"`
	UpdatedAt time.Time            `json:"updated_at" example:"2022-10-07T06:49:42.629210158+07:00"`
}
