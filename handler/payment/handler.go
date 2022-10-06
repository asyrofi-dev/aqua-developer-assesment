package handler

import usecase "aqua-developer-assesment/usecase/payment"

type PaymentHandler struct {
	PaymentUC usecase.IPaymentUsecase
}

func NewPaymentHandler(usecase usecase.IPaymentUsecase) *PaymentHandler {
	return &PaymentHandler{usecase}
}
