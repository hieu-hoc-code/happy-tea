package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"server/controllers"
	"server/middlewares"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Init() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Welcome).Methods("GET")

	// auth user
	router.HandleFunc("/api/register", controllers.Register).Methods("POST")
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/logout", controllers.Logout).Methods("GET")

	// auth admin
	router.HandleFunc("/api/login/admin", controllers.LoginAdmin).Methods("POST")
	router.HandleFunc("/api/register/admin", controllers.RegisterAdmin).Methods("POST")
	router.HandleFunc("/api/admin", middlewares.IsAuthenticated(controllers.GetAdmin)).Methods("GET")

	// middleware
	IsAuthenticated := middlewares.IsAuthenticated
	router.HandleFunc("/api/user", IsAuthenticated(controllers.VerifyCache(controllers.User))).Methods("GET")
	router.HandleFunc("/api/user", controllers.UpdateUser).Methods("POST")

	// product
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", controllers.ProductCache(controllers.AllProducts)).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/api/products", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

	// search with ES
	router.HandleFunc("/api/search", controllers.ProductCache(controllers.SearchProduct)).Queries("q", "{q}").Methods("GET")

	// cart
	router.HandleFunc("/api/cart", IsAuthenticated(controllers.CreateCart)).Methods("POST")
	router.HandleFunc("/api/cartitems/{id}", IsAuthenticated(controllers.GetCartItem)).Methods("GET")
	router.HandleFunc("/api/cartitems", controllers.CreateCartItem).Methods("POST")
	router.HandleFunc("/api/cartitems/{id}", controllers.DeleteCartItem).Methods("DELETE")

	// user_payment
	router.HandleFunc("/api/userpayments", IsAuthenticated(controllers.CreateUserPayment)).Methods("POST")

	// order
	router.HandleFunc("/api/orders", IsAuthenticated(controllers.CreateOrder)).Methods("POST")
	router.HandleFunc("/api/orders", IsAuthenticated(controllers.GetAllOrders)).Methods("GET")
	//
	router.HandleFunc("/api/order-detail/{id}", controllers.GetOrderDetail).Methods("GET")

	// catalog
	router.HandleFunc("/api/catalog", IsAuthenticated(controllers.CreateCatalog)).Methods("POST")
	router.HandleFunc("/api/catalog", controllers.GetCatalog).Methods("GET")
	router.HandleFunc("/api/catalog", IsAuthenticated(controllers.UpdateCatalog)).Methods("PUT")
	router.HandleFunc("/api/catalog/{id}", IsAuthenticated(controllers.DeleteCatalog)).Methods("DELETE")

	//brand
	router.HandleFunc("/api/brands", IsAuthenticated(controllers.CreateBrand)).Methods("POST")
	router.HandleFunc("/api/brands", controllers.GetBrand).Methods("GET")

	//topping
	router.HandleFunc("/api/toppings", IsAuthenticated(controllers.CreateTopping)).Methods("POST")
	router.HandleFunc("/api/toppings", controllers.GetTopping).Methods("GET")

	// ratting
	router.HandleFunc("/api/ratings", IsAuthenticated(controllers.CreateRating)).Methods("POST")
	router.HandleFunc("api/ratings", controllers.GetRating).Methods("GET")

	// address
	router.HandleFunc("/api/address", controllers.CreateAddress).Methods("POST")
	router.HandleFunc("/api/address", controllers.GetAddress).Methods("GET")
	router.HandleFunc("/api/address", controllers.UpdateAddress).Methods("PUT")
	router.HandleFunc("/api/address", controllers.DeleteAddress).Queries("id", "{id}", "user_id", "{user_id}").Methods("DELETE")

	// user
	router.HandleFunc("/api/upload/user", controllers.UploadImageUser).Methods("PATCH")

	s := http.StripPrefix("/static", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static").Handler(s)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	fmt.Println("server is running http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}

func StopRouter(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting router")
			wg.Done()
			return
		}
	}
}
