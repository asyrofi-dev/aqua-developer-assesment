package repository

import "aqua-developer-assesment/entity"

func (repo ProductRepository) ListProducts() ([]entity.Product, error) {
	var result []entity.Product
	query := repo.DB.Debug().Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
