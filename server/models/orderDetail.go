package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	Id        uint    `json:"id" gorm:"primaryKey"`
	ProductId uint    `json:"product_id"`
	Name      string  `json:"name"`
	Quantity  uint    `json:"quantity"`
	OrderId   uint    `json:"order_id"`
	Price     float64 `json:"price"`
}
