package ordersService

import (
	"errors"
	"mini-indobat/models"
	"mini-indobat/service/productsService"
	"time"

	"gorm.io/gorm"
)

type OrdersService struct {
	ordersRepo   OrdersRepo
	productsRepo productsService.ProductsRepo
}

type Service interface {
	CreateOrder(data models.Orders) (models.Orders, error)
}

func NewOrdersService(ordersRepo OrdersRepo, productsRepo productsService.ProductsRepo) Service {
	return &OrdersService{
		ordersRepo:   ordersRepo,
		productsRepo: productsRepo,
	}
}

func (s *OrdersService) CreateOrder(data models.Orders) (models.Orders, error) {
	var result models.Orders

	err := s.ordersRepo.WithTransaction(func(tx *gorm.DB) error {
		product, err := s.productsRepo.GetProductForUpdate(tx, data.ProductId)
		if err != nil {
			return err
		}

		if *product.Stok < data.Quantity {
			return errors.New("stok tidak cukup")
		}

		*product.Stok -= data.Quantity

		if err := s.productsRepo.UpdateProductTx(tx, product); err != nil {
			return err
		}

		data.Subtotal = product.Harga * float64(data.Quantity)
		discountPrice := data.Subtotal * (data.DiscountPercent / 100)
		data.Total = data.Subtotal - discountPrice
		data.CreatedAt = time.Now()
		result, err = s.ordersRepo.CreateOrderTx(tx, data)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return models.Orders{}, err
	}

	return result, nil
}
