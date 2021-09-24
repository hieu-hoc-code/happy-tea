package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Id         uint `json:"ratting" gorm:"Primary_key"`
	User_ID    uint `json:"userId" gorm:"foreign_key"`
	Rate       int  `json:"rate"`
	Product_ID uint `json:"productId" gorm:"foreign_key"`
}
