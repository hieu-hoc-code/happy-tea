package database

import (
	"context"
	"log"
	"server/config"
	"server/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var Ctx = context.Background()

var Cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func Connect() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	database, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database!")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Product{}, &models.Catalog{}, &models.Cart{}, &models.CartItem{}, &models.Order{}, &models.OrderDetail{}, &models.UserPayment{}, &models.Topping{}, &models.Brand{}, &models.Rating{}, &models.Admin{}, &models.Address{})
}
