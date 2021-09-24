package models

type Cache struct {
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Token       string `json:"token"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Sex         string `json:"sex"`
	Birthday    string `json:"birthday"`
	Image       string `json:"image"`
	CartID      string `json:"cart_id"`
}
