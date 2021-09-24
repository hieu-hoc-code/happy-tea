package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Id          uint       `json:"id" gorm:"primaryKey"`
	Name        string     `json:"name"`
	Email       string     `json:"email" gorm:"unique"`
	Password    []byte     `json:"-"`
	PhoneNumber string     `json:"phone_number"`
	Sex         string     `json:"sex"`
	Birthday    *time.Time `json:"birthday"`
	Admin       bool       `json:"active"`
}
