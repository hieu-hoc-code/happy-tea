package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Id   uint   `json:"id" gorm:"Primary_key"`
	Name string `json:"name"`
}
