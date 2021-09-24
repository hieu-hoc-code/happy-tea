package controllers

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"

	"github.com/sirupsen/logrus"
)

func GetRating(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get Rating")
	var ratings []models.Rating

	database.DB.Find(&ratings)

	if err := json.NewEncoder(w).Encode(ratings); err != nil {
		logrus.Error("encode fail :", err)
		json.NewEncoder(w).Encode("fetch data fail")
		return
	}
}

func CreateRating(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Create Rating")
	var rating models.Rating
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&rating)
	if err != nil {
		logrus.Error("parse data fail: ", err)
		json.NewEncoder(w).Encode("Parse data fail")
		return
	}

	database.DB.Where("id = ?", rating.Product_ID).First(&product)

	//update product rating
	product.NumOfRate++
	temp := product.Rating*float32(product.NumOfRate-1) + float32(rating.Rate)
	product.Rating = temp / float32(product.NumOfRate)

	database.DB.Save(&product)
	database.DB.Create(&rating)
	json.NewEncoder(w).Encode("create rating success")
}
