package controllers

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"

	"github.com/sirupsen/logrus"
)

func GetBrand(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get Brand")
	var brands []models.Brand

	database.DB.Find(&brands)

	if err := json.NewEncoder(w).Encode(brands); err != nil {
		logrus.Error("encode fail :", err)
		json.NewEncoder(w).Encode("fetch data fail")
		return
	}
}

func CreateBrand(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Create Brand")
	var brand models.Brand

	err := json.NewDecoder(r.Body).Decode(&brand)
	if err != nil {
		logrus.Error("parse data fail: ", err)
		json.NewEncoder(w).Encode("Parse data fail")
		return
	}

	database.DB.Create(&brand)
}
