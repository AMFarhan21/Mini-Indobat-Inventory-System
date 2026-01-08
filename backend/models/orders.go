package models

import (
	"time"
)

type Orders struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductId       int       `json:"product_id"`
	Quantity        int       `json:"quantity"`
	Subtotal        float64   `json:"subtotal"`
	DiscountPercent float64   `json:"discount_percent"`
	Total           float64   `json:"total"`
	CreatedAt       time.Time `json:"created_at"`
}
