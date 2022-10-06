package usecase

import (
	"aqua-developer-assesment/entity"

	"github.com/jinzhu/copier"
)

func (usecase ProductUsecase) GetProduct(id int) (entity.ProductResponse, error) {
	var response entity.ProductResponse

	result, err := usecase.ProductRepo.GetProduct(id)

	if err != nil {
		return entity.ProductResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
