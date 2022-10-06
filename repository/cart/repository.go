package repository

import (
	"aqua-developer-assesment/entity"

	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(cart entity.Cart) (entity.Cart, error)
	GetCart(id int) (entity.Cart, error)
}

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db}
}
