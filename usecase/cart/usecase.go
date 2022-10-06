package usecase

import (
	"aqua-developer-assesment/entity"
	repository "aqua-developer-assesment/repository/cart"
)

type ICartUsecase interface {
	CreateCart(cart entity.CartRequest) (entity.CartResponse, error)
	GetCart(id int) (entity.CartResponse, error)
}

type CartUsecase struct {
	CartRepo repository.ICartRepository
}

func NewCartUsecase(repo repository.ICartRepository) *CartUsecase {
	return &CartUsecase{repo}
}
