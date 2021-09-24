package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/database"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateCartItem(w http.ResponseWriter, r *http.Request) {
	// parse body
	var data models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	// check have product_id yet, update when have
	var cartItem models.CartItem
	database.DB.Where("product_id = ?", data.ProductID).First(&cartItem)
	if cartItem.ID != 0 {
		cartItem.Quantity = data.Quantity
		database.DB.Model(&cartItem).Updates(cartItem)

		fmt.Fprintf(w, "updated cartItem")
		return
	}

	// if have not product_id
	if err := database.DB.Create(&data); err == nil {
		fmt.Fprintf(w, "failed create cartItem")
		return
	}

	// add len in cart
	var cart models.Cart
	database.DB.Where("id = ?", data.CartID).First(&cart)
	cart.LenItem = cart.LenItem + 1
	database.DB.Model(&cart).Updates(cart)

	fmt.Fprintf(w, "%v", data)
}

func GetCartItem(w http.ResponseWriter, r *http.Request) {
	// get cart_id
	vars := mux.Vars(r)
	cartID, _ := strconv.Atoi(vars["id"])
	type cartItem struct {
		ID        uint    `json:"id"`
		ProductID int     `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Name      string  `json:"name"`
		Descrip   string  `json:"descrip"`
		Price     float64 `json:"price"`
		Image     string  `json:"image"`
	}

	var data []cartItem
	database.DB.Table("cart_items").Select("cart_items.id, product_id, quantity, name, descrip, price, image").Joins("LEFT JOIN products ON cart_items.product_id = products.id").Where("cart_id = ?", cartID).Find(&data)
	json.NewEncoder(w).Encode(&data)
}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartID, _ := strconv.Atoi(vars["id"])
	var cartItem models.CartItem

	database.DB.Delete(&cartItem, cartID)
	w.WriteHeader(http.StatusOK)
}
