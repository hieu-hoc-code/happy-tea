package models

import "gorm.io/gorm"

type Topping struct {
	gorm.Model
	Id    uint    `json:"id" gorm:"Primary_key"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
