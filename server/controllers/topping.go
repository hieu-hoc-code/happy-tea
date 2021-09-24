package controllers

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"

	"github.com/sirupsen/logrus"
)

func GetTopping(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get Topping")
	var toppings []models.Topping

	database.DB.Find(&toppings)

	if err := json.NewEncoder(w).Encode(toppings); err != nil {
		logrus.Error("encode fail :", err)
		json.NewEncoder(w).Encode("fetch data fail")
		return
	}
}

func CreateTopping(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Create Topping")
	var topping models.Topping

	err := json.NewDecoder(r.Body).Decode(&topping)
	if err != nil {
		logrus.Error("parse data fail: ", err)
		json.NewEncoder(w).Encode("Parse data fail")
		return
	}

	database.DB.Create(&topping)
}
