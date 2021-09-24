package controllers

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"
	"strconv"
)

func CreateCart(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		response := map[string]string{
			"msg": "cannot get user_id",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	userID, _ := strconv.Atoi(cookie.Value)
	// check user have cart yet
	var cart models.Cart
	database.DB.Where("user_id = ?", userID).First(&cart)
	if cart.ID != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&cart.ID)
		return
	}

	cart.UserID = uint(userID)
	database.DB.Create(&cart)
	json.NewEncoder(w).Encode(&cart.ID)
}
