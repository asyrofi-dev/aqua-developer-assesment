package usecase

import (
	"aqua-developer-assesment/entity"

	"github.com/jinzhu/copier"
)

func (usecase CartUsecase) CreateCart(cart entity.CartRequest) (entity.CartResponse, error) {
	var newCart entity.Cart
	var response entity.CartResponse

	copier.Copy(&newCart, &cart)

	result, err := usecase.CartRepo.CreateCart(newCart)

	if err != nil {
		return entity.CartResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
