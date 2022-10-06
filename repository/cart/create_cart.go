package repository

import "aqua-developer-assesment/entity"

func (repo CartRepository) CreateCart(cart entity.Cart) (entity.Cart, error) {
	query := repo.DB.Debug().Create(&cart)

	if query.Error != nil {
		return entity.Cart{}, query.Error
	}

	query = repo.DB.Debug().Preload("Items.Product").First(&cart)

	if query.Error != nil {
		return entity.Cart{}, query.Error
	}

	return cart, nil
}
