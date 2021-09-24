package models

import (
	"gorm.io/gorm"
)

type Catalog struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
