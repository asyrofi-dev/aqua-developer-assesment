package usecase

import (
	"aqua-developer-assesment/entity"
	repository "aqua-developer-assesment/repository/product"
)

type IProductUsecase interface {
	CreateProduct(product entity.ProductRequest) (entity.ProductResponse, error)
	ListBooks() ([]entity.ProductResponse, error)
	GetProduct(id int) (entity.ProductResponse, error)
}

type ProductUsecase struct {
	ProductRepo repository.IProductRepository
}

func NewProductUsecase(repo repository.IProductRepository) *ProductUsecase {
	return &ProductUsecase{repo}
}
