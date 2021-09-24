package controllers

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetCatalog(w http.ResponseWriter, r *http.Request) {
	var catalogs []models.Catalog

	database.DB.Find(&catalogs)

	if err := json.NewEncoder(w).Encode(catalogs); err != nil {
		logrus.Error("encode fail :", err)
		json.NewEncoder(w).Encode("fetch data fail")
		return
	}
}

func CreateCatalog(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Create catalog")
	var catalog models.Catalog

	err := json.NewDecoder(r.Body).Decode(&catalog)
	if err != nil {
		logrus.Error("parse data fail: ", err)
		json.NewEncoder(w).Encode("Parse data fail")
		return
	}

	database.DB.Create(&catalog)
	json.NewEncoder(w).Encode(&catalog)
}

func UpdateCatalog(w http.ResponseWriter, r *http.Request) {
	logrus.Info("update catalog")
	var catalog models.Catalog

	if err := json.NewDecoder(r.Body).Decode(&catalog); err != nil {
		logrus.Error("error when parse body: ", err)
		json.NewEncoder(w).Encode("parse body fail")
		return
	}
	database.DB.Save(&catalog)
}

func DeleteCatalog(w http.ResponseWriter, r *http.Request) {
	logrus.Info("delete catalog")
	var catalog models.Catalog
	id := mux.Vars(r)["id"]
	i, _ := strconv.Atoi(id)

	database.DB.Delete(&catalog, i)
}
