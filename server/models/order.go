package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id                uint      `json:"id" gorm:"primaryKey"`
	UserId            uint      `json:"user_id"`
	Total             float64   `json:"total"`
	PaymentId         uint      `json:"payment_id"`
	AddressId         uint      `json:"address_id"`
	CustomerName      string    `json:"customer_name"`
	Email             string    `json:"email"`
	ThankYouEmailSent bool      `json:"thankyou_email_sent"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
	OrderDetailId     uint      `json:"orderDetailId"`
}
