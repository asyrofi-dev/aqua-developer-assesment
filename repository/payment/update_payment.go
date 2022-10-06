package repository

import "aqua-developer-assesment/entity"

func (repo PaymentRepository) UpdatePayment(Payment entity.Payment) (entity.Payment, error) {
	query := repo.DB.Debug().Model(&Payment).Updates(Payment)

	if query.Error != nil {
		return entity.Payment{}, query.Error
	}

	return Payment, nil
}
