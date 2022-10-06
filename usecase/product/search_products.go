package usecase

import (
	"aqua-developer-assesment/entity"

	"github.com/jinzhu/copier"
)

func (usecase ProductUsecase) SearchProducts(filter entity.ProductFilter) ([]entity.ProductResponse, error) {
	var response []entity.ProductResponse

	result, err := usecase.ProductRepo.SearchProducts(filter)

	if err != nil {
		return nil, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
