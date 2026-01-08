package productsService

import (
	"context"
	"mini-indobat/models"

	"gorm.io/gorm"
)

type ProductsRepo interface {
	GetAllProducts(ctx context.Context) ([]models.Products, error)
	GetProductForUpdate(tx *gorm.DB, id int) (models.Products, error)
	CreateProduct(ctx context.Context, data models.Products) (models.Products, error)
	UpdateProductTx(tx *gorm.DB, data models.Products) error
	DeleteProduct(ctx context.Context, id int) error
}
