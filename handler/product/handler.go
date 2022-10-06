package handler

import usecase "aqua-developer-assesment/usecase/product"

type ProductHandler struct {
	ProductUC usecase.IProductUsecase
}

func NewProductHandler(usecase usecase.IProductUsecase) *ProductHandler {
	return &ProductHandler{usecase}
}
