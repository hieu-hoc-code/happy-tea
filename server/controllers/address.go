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

func CreateAddress(w http.ResponseWriter, r *http.Request) {
	var data models.Address
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", err)
		return
	}
	database.DB.Create(&data)

	json.NewEncoder(w).Encode(data)
}

func GetAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var data []models.Address

	database.DB.Where("user_id = ?", userID).Find(&data)

	json.NewEncoder(w).Encode(data)
}

func UpdateAddress(w http.ResponseWriter, r *http.Request) {

	var data models.Address
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", err)
		return
	}

	database.DB.Model(&data).Updates(data)

	json.NewEncoder(w).Encode(data)
}

func DeleteAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	var data []models.Address

	database.DB.Where("user_id = ?", userID).Find(&data)
	if len(data) < 2 {
		fmt.Fprintf(w, "khong the xoa")
		return
	}

	id := vars["id"]

	idUint, _ := strconv.ParseUint(id, 10, 64)
	address := models.Address{
		ID: uint(idUint),
	}
	database.DB.Delete(&address)

	fmt.Fprintf(w, "deleted")
}
