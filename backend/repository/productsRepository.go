package repository

import (
	"context"
	"errors"
	"mini-indobat/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductsRepository struct {
	DB *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		DB: db,
	}
}

func (r *ProductsRepository) GetAllProducts(ctx context.Context) ([]models.Products, error) {
	var products []models.Products
	err := r.DB.Table("products").WithContext(ctx).Order("id DESC").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductsRepository) GetProductForUpdate(tx *gorm.DB, id int) (models.Products, error) {
	var product models.Products

	err := tx.Model(&models.Products{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("id=?", id).First(&product).Error
	if err != nil {
		return models.Products{}, err
	}

	return product, nil
}

func (r *ProductsRepository) CreateProduct(ctx context.Context, data models.Products) (models.Products, error) {
	err := r.DB.Table("products").WithContext(ctx).Create(&data).Error
	if err != nil {
		return models.Products{}, err
	}

	return data, nil
}

func (r *ProductsRepository) UpdateProductTx(tx *gorm.DB, data models.Products) error {
	row := tx.Model(&models.Products{}).Where("id=?", data.Id).Updates(&data)
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (r *ProductsRepository) DeleteProduct(ctx context.Context, id int) error {
	row := r.DB.Table("products").WithContext(ctx).Where("id=?", id).Delete(&models.Products{})
	if err := row.Error; err != nil {
		return err
	}

	if row.RowsAffected == 0 {
		return errors.New("product not found")
	}

	return nil
}
