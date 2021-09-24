package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id           uint    `json:"id"`
	Name         string  `json:"name"`
	Descrip      string  `json:"desc"`
	Price        float64 `json:"price"`
	NumOfRate    uint    `json:"numOfRate"`
	Rating       float32 `json:"rating"`
	Image        string  `json:"image"`
	Brand_ID     int     `json:"brandId" gorm:"foreign_key"`
	Topping_ID   int     `json:"toppingId" gorm:"foreign_key"`
	Category_ID  int     `json:"categoryId" gorm:"foreign_key"`
	Inventory_ID int     `json:"inventoryId" gorm:"foreign_key"`
	Discount_ID  int     `json:"discountId" gorm:"foreign_key"`
	Shipping     bool    `json:"shipping"`
}
