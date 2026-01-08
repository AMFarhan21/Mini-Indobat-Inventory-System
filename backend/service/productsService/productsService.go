package productsService

import (
	"context"
	"mini-indobat/models"
)

type ProductsService struct {
	productsRepo ProductsRepo
}

type Service interface {
	GetAllProducts(ctx context.Context) ([]models.Products, error)
	CreateProduct(ctx context.Context, data models.Products) (models.Products, error)
}

func NewProductsService(productsRepo ProductsRepo) Service {
	return &ProductsService{
		productsRepo: productsRepo,
	}
}

func (s *ProductsService) GetAllProducts(ctx context.Context) ([]models.Products, error) {
	return s.productsRepo.GetAllProducts(ctx)
}

func (s *ProductsService) CreateProduct(ctx context.Context, data models.Products) (models.Products, error) {
	return s.productsRepo.CreateProduct(ctx, data)
}
