package repository

import "aqua-developer-assesment/entity"

func (repo ProductRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	query := repo.DB.Debug().Model(&product).Updates(product)

	if query.Error != nil {
		return entity.Product{}, query.Error
	}

	return product, nil
}
