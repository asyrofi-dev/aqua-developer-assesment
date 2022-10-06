package repository

import (
	"aqua-developer-assesment/entity"
)

func (repo PaymentRepository) CreatePayment(payment entity.Payment) (entity.Payment, error) {
	query := repo.DB.Debug().Create(&payment)

	if query.Error != nil {
		return entity.Payment{}, query.Error
	}

	query = repo.DB.Debug().Preload("Cart.Items.Product").First(&payment)

	if query.Error != nil {
		return entity.Payment{}, query.Error
	}

	return payment, nil
}
