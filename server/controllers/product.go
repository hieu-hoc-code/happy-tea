package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"server/database"
	"server/models"
	"server/util"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("create product")

	file, _, err := r.FormFile("image")
	if err != nil {
		logrus.Error("read form fail: ", err)
		json.NewEncoder(w).Encode(err)
		return
	}

	defer file.Close()

	tempFile, err := ioutil.TempFile("static/product", "product-*.jpg")
	if err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Info(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	tempFile.Write(fileBytes)
	url := "http://localhost:3000/" + tempFile.Name()

	var data models.Product
	if err := json.Unmarshal([]byte(r.FormValue("newProduct")), &data); err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	data.Image = url
	database.DB.Create(&data)

	// create es
	conn := models.ConnectES()
	es := models.NewES(conn)
	es.Create(data)
	json.NewEncoder(w).Encode(data)
	//clear cache
	ClearProductCache(*database.Cache, database.Ctx)

}

func formatResquestForm(i string, v []string) string {
	switch i {
	case "categoryId[]":
		s := strings.Join(v, ",")
		t := "(" + s + ")"
		return fmt.Sprintf("category_id in %s", t)
	case "priceRange":
		price, _ := strconv.ParseFloat(v[0], 32)
		return fmt.Sprintf("price <= %v", price/1000)
	case "rating":
		return fmt.Sprintf("rating >= %v", v[0])
	case "brandId[]":
		s := strings.Join(v, ",")
		t := "(" + s + ")"
		return fmt.Sprintf("brand_id in %s", t)
	case "toppingId[]":
		s := strings.Join(v, ",")
		t := "(" + s + ")"
		return fmt.Sprintf("topping_id in %s", t)
	case "shipping":
		return fmt.Sprintf("shipping = %s", v[0])
	default:
		return ""
	}
}

func AllProducts(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Get product by query")
	var products []models.Product
	r.ParseForm()
	query := []string{}
	for i, v := range r.Form {
		if v[0] != "" {
			query = append(query, formatResquestForm(i, v))
		}
	}
	str := strings.Join(query, " AND ")
	database.DB.Where(str).Find(&products)
	json.NewEncoder(w).Encode(products)

	//cache products by query
	data, err := json.Marshal(&products)
	if err != nil {
		logrus.Error(err)
	}
	url := r.URL.String()
	fmt.Println("ben trong: ")

	cacheErr := database.Cache.Set(database.Ctx, url, data, 60*time.Second).Err()
	if cacheErr != nil {
		logrus.Error(err, nil)
	}
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Error("string convert fail :", err)
		return
	}
	var data models.Product

	database.DB.Where("id = ?", uint(id)).First(&data)
	if data.Id == 0 {
		fmt.Fprintf(w, "not found product_id: %v", id)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("update product")
	//load anh
	file, _, err := r.FormFile("image")
	if err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("static/product", "product-*.jpg")
	if err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	defer tempFile.Close()

	fileByte, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	tempFile.Write(fileByte)

	url := "http://localhost:3000/" + tempFile.Name()

	// parse data
	var data models.Product
	if err := json.Unmarshal([]byte(r.FormValue("newProduct")), &data); err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode(err)
		return
	}

	if data.Image != "" {
		filePath := data.Image[21:]
		filePath = "." + filePath
		util.RemoveFile(filePath)
	}

	data.Image = url

	database.DB.Model(&data).Updates(data)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logrus.Info("send response false")
		json.NewEncoder(w).Encode("send response false")
		return
	}
	// create es
	conn := models.ConnectES()
	es := models.NewES(conn)
	es.Update(data)
	//clear cache
	ClearProductCache(*database.Cache, database.Ctx)

	json.NewEncoder(w).Encode(data)
	database.Cache.FlushDB(database.Ctx)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error("string convert fail :", err)
		return
	}
	data := models.Product{
		Id: uint(i),
	}
	database.DB.Delete(&data)
	// create es
	conn := models.ConnectES()
	es := models.NewES(conn)
	es.Delete(data.Id)

	//clear cache
	ClearProductCache(*database.Cache, database.Ctx)

	fmt.Fprintf(w, "deleted product_id: %v", id)
}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Search by ES")

	vars := mux.Vars(r)
	query := vars["q"]
	// search product with ElasticSearch
	conn := models.ConnectES()
	es := models.NewES(conn)
	data := es.Search(query)

	json.NewEncoder(w).Encode(data)
	//cache query
	/* 	products, err := json.Marshal(&data)
	   	if err != nil {
	   		logrus.Error(err)
	   	}
	   	url := r.URL.String()
	   	fmt.Println("ben trong: ")
	   	cacheErr := database.Cache.Set(database.Ctx, url, products, 60*time.Second).Err()
	   	if cacheErr != nil {
	   		logrus.Error(cacheErr)
	   	} */
}
