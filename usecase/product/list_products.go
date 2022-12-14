package usecase

import (
	"aqua-developer-assesment/entity"

	"github.com/jinzhu/copier"
)

func (usecase ProductUsecase) ListProducts() ([]entity.ProductResponse, error) {
	var response []entity.ProductResponse

	result, err := usecase.ProductRepo.ListProducts()

	if err != nil {
		return nil, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
