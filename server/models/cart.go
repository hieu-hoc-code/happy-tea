package models

type Cart struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	UserID  uint `json:"user_id" gorm:"foreignKey"`
	LenItem int  `json:"len_item" gorm:"default:0"`
}
