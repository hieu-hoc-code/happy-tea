package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"server/database"
	"server/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Product struct {
	Name      string  `json:"name"`
	ProductId uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Price     float64 `json:"price"`
}

type DataOrder struct {
	Total     float64   `json:"total"`
	PaymentID uint      `json:"payment_id"`
	AddressID uint      `json:"address_id"`
	Products  []Product `json:"products"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	database.DB.Transaction(func(tx *gorm.DB) error {
		cookie, _ := r.Cookie("user_id")

		userID, _ := strconv.Atoi(cookie.Value)

		// parser
		var data DataOrder
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			fmt.Fprintf(w, "%v", err)
			return nil
		}

		// lay email
		var user models.User
		database.DB.Where("id = ?", userID).First(&user)
		if user.Id == 0 {
			fmt.Fprintf(w, "not found email %v", data)
			return nil
		}

		order := models.Order{
			UserId:            uint(userID),
			Total:             data.Total,
			PaymentId:         data.PaymentID,
			AddressId:         data.AddressID,
			Email:             user.Email,
			CustomerName:      user.Name,
			ThankYouEmailSent: false,
			CreatedAt:         time.Now(),
			ModifiedAt:        time.Now(),
		}

		if err := tx.Create(&order).Error; err != nil {
			fmt.Fprintf(w, "cannot create order")
			return err
		}

		for _, v := range data.Products {
			order_detail := models.OrderDetail{
				ProductId: v.ProductId,
				Quantity:  v.Quantity,
				OrderId:   order.Id,
				Price:     v.Price,
				Name:      v.Name,
			}
			fmt.Println(v, v.Name)
			if err := tx.Create(&order_detail).Error; err != nil {
				fmt.Fprintf(w, "cannot create order detail")
				return err
			}
		}

		fmt.Fprintf(w, "order thanh cong: %v", order)
		return nil
	})
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	// check dang nhap
	cookie, err := r.Cookie("user_id")
	if err != nil {
		fmt.Fprintf(w, "Tai khoan chua dang nhap %v", err)
		return
	}

	user_id, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Cannot parse string to uint: %v", err)
	}

	var orders []models.Order
	database.DB.Where("user_id = ?", user_id).Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrderDetail(w http.ResponseWriter, r *http.Request) {
	logrus.Info("get order detail")
	id := mux.Vars(r)["id"]
	i, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode("Can't get  order detail id")
		return
	}
	var data []models.OrderDetail
	database.DB.Where("order_id = ?", uint(i)).Find(&data)

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		logrus.Error(err)
		json.NewEncoder(w).Encode("send data fail")
	}
}
