package models

import "gorm.io/gorm"

type UserPayment struct {
	gorm.Model
	Id     uint   `json:"id" gorm:"primaryKey"`
	Type   string `json:"type"`
	UserId uint   `json:"user_id" gorm:"foreignKey"`
}
