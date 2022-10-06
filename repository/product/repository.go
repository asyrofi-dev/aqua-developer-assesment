package repository

import (
	"aqua-developer-assesment/entity"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	ListProducts() ([]entity.Product, error)
	GetProduct(id int) (entity.Product, error)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}
