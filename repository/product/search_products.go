package repository

import "aqua-developer-assesment/entity"

func (repo ProductRepository) SearchProducts(filter entity.ProductFilter) ([]entity.Product, error) {
	var result []entity.Product

	query := repo.DB.Debug().Where("category LIKE ? AND price >= ? AND price <= ?",
		filter.Category, filter.MinPrice, filter.MaxPrice).Find(&result)

	if query.Error != nil {
		return nil, query.Error
	}

	return result, nil
}
