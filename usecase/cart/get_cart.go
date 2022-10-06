package usecase

import (
	"aqua-developer-assesment/entity"

	"github.com/jinzhu/copier"
)

func (usecase CartUsecase) GetCart(id int) (entity.CartResponse, error) {
	var response entity.CartResponse

	result, err := usecase.CartRepo.GetCart(id)

	if err != nil {
		return entity.CartResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
