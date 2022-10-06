package usecase

import (
	"aqua-developer-assesment/entity"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jinzhu/copier"
)

func (usecase PaymentUsecase) CreatePayment(payment entity.PaymentRequest) (entity.PaymentResponse, error) {
	var newPayment entity.Payment
	var response entity.PaymentResponse

	copier.Copy(&newPayment, &payment)

	result, err := usecase.PaymentRepo.CreatePayment(newPayment)

	if err != nil {
		return entity.PaymentResponse{}, err
	}

	// Upload Photo

	src, err := payment.Photo.Open()

	if err != nil {
		return entity.PaymentResponse{}, err
	}

	defer src.Close()

	filename := fmt.Sprintf("payment_%d%s", result.ID, filepath.Ext(payment.Photo.Filename))

	path := filepath.Join("public", "images", "payments", filename)

	dst, err := os.Create(path)

	if err != nil {
		return entity.PaymentResponse{}, err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return entity.PaymentResponse{}, err
	}

	// update result

	result.Photo = fmt.Sprintf("/api/v1/images/payments/%s", filename)

	result, err = usecase.PaymentRepo.UpdatePayment(result)

	if err != nil {
		return entity.PaymentResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
