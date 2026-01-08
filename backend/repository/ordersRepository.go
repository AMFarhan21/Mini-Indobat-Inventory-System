package repository

import (
	"mini-indobat/models"

	"gorm.io/gorm"
)

type OrdersRepository struct {
	DB *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) *OrdersRepository {
	return &OrdersRepository{
		DB: db,
	}
}

func (r *OrdersRepository) CreateOrderTx(tx *gorm.DB, data models.Orders) (models.Orders, error) {

	err := tx.Model(&models.Orders{}).Create(&data).Error
	if err != nil {
		return models.Orders{}, err
	}

	return data, nil
}

func (r *OrdersRepository) WithTransaction(fn func(tx *gorm.DB) error) error {
	return r.DB.Model(&models.Orders{}).Transaction(fn)
}
