package repository

import (
	"aqua-developer-assesment/entity"

	"gorm.io/gorm"
)

type IPaymentRepository interface {
	CreatePayment(Payment entity.Payment) (entity.Payment, error)
	UpdatePayment(Payment entity.Payment) (entity.Payment, error)
	GetPayment(id int) (entity.Payment, error)
}

type PaymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db}
}
