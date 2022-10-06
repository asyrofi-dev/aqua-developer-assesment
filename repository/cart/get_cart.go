package repository

import "aqua-developer-assesment/entity"

func (repo CartRepository) GetCart(id int) (entity.Cart, error) {
	var result entity.Cart
	query := repo.DB.Debug().Preload("Items.Product").First(&result)

	if query.Error != nil {
		return entity.Cart{}, query.Error
	}

	return result, nil
}
