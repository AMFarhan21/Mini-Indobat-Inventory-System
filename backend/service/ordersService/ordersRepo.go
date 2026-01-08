package ordersService

import (
	"mini-indobat/models"

	"gorm.io/gorm"
)

type OrdersRepo interface {
	WithTransaction(fn func(tx *gorm.DB) error) error
	CreateOrderTx(tx *gorm.DB, data models.Orders) (models.Orders, error)
}
