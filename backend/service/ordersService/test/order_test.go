package test

import (
	"mini-indobat/models"
	"mini-indobat/repository"
	"mini-indobat/service/ordersService"
	"mini-indobat/utils/database"
	"sync"
	"sync/atomic"
	"testing"
)

func TestConcurrentOrder(t *testing.T) {
	// config := config.Load()
	db := database.GetTestDatabaseConnection("postgresql://postgres@localhost:5432/mini_indobat_test?sslmode=disable")
	productsRepo := repository.NewProductsRepository(db)
	ordersRepo := repository.NewOrdersRepository(db)
	ordersService := ordersService.NewOrdersService(ordersRepo, productsRepo)

	var wg sync.WaitGroup
	success := int32(0)
	failed := int32(0)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := ordersService.CreateOrder(models.Orders{
				ProductId: 1,
				Quantity:  1,
			})
			if err != nil {
				atomic.AddInt32(&failed, 1)
			} else {
				atomic.AddInt32(&success, 1)
			}
		}()
	}

	wg.Wait()

	if success != 1 {
		t.Errorf("expected 1 success, got %d", success)
	}

	if failed != 9999 {
		t.Errorf("expected 9999 failed, got %d", failed)
	}

}
