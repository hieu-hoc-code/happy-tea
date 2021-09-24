package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"Primary_key"`
	UserID      uint   `json:"user_id" gorm:"foreign_key"`
	Name        string `json:"name"`
	Company     string `json:"company"`
	PhoneNumber string `json:"phone_number"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Village     string `json:"village"`
	Address     string `json:"address"`
	TypeAddress string `json:"type_address"`
	Official    bool   `json:"official"`
}
