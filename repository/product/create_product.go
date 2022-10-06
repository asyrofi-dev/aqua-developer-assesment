package repository

import (
	"aqua-developer-assesment/entity"
)

func (repo ProductRepository) CreateProduct(product entity.Product) (entity.Product, error) {
	query := repo.DB.Debug().Create(&product)

	if query.Error != nil {
		return entity.Product{}, query.Error
	}

	return product, nil
}
