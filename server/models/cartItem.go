package models

type CartItem struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CartID    int    `json:"cart_id"`
}
