package handler

import usecase "aqua-developer-assesment/usecase/cart"

type CartHandler struct {
	CartUC usecase.ICartUsecase
}

func NewCartHandler(usecase usecase.ICartUsecase) *CartHandler {
	return &CartHandler{usecase}
}
