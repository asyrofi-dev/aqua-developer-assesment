package usecase

import (
	"aqua-developer-assesment/entity"

	"github.com/jinzhu/copier"
)

func (usecase PaymentUsecase) GetPayment(id int) (entity.PaymentResponse, error) {
	var response entity.PaymentResponse

	result, err := usecase.PaymentRepo.GetPayment(id)

	if err != nil {
		return entity.PaymentResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
