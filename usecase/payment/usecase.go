package usecase

import (
	"aqua-developer-assesment/entity"
	repository "aqua-developer-assesment/repository/payment"
)

type IPaymentUsecase interface {
	CreatePayment(payment entity.PaymentRequest) (entity.PaymentResponse, error)
	GetPayment(id int) (entity.PaymentResponse, error)
}

type PaymentUsecase struct {
	PaymentRepo repository.IPaymentRepository
}

func NewPaymentUsecase(repo repository.IPaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{repo}
}
