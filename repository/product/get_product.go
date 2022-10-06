package repository

import "aqua-developer-assesment/entity"

func (repo ProductRepository) GetProduct(id int) (entity.Product, error) {
	var result entity.Product
	query := repo.DB.Debug().First(&result, id)

	if query.Error != nil {
		return entity.Product{}, query.Error
	}

	return result, nil
}
