package repository

import "aqua-developer-assesment/entity"

func (repo PaymentRepository) GetPayment(id int) (entity.Payment, error) {
	var result entity.Payment
	query := repo.DB.Debug().Preload("Cart.Items.Product").First(&result, id)

	if query.Error != nil {
		return entity.Payment{}, query.Error
	}

	return result, nil
}
