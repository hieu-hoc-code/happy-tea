package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"server/database"
	"server/models"
	"server/util"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// parse body
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	// check password & confirm password
	if data["password"] != data["confirm"] {
		fmt.Fprintf(w, "password do not match")
		return
	}

	// check email existed
	var isExist models.User
	database.DB.Where("email = ?", data["email"]).First(&isExist)
	if isExist.Email != "" {
		fmt.Fprintf(w, "email exists")
		return
	}

	// hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	var user models.User
	user.Name = data["name"]
	user.Email = data["email"]
	user.Password = password

	database.DB.Create(&user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Đăng ký tài khoản thành công <3")
}

func Login(w http.ResponseWriter, r *http.Request) {
	// parse body
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errors := []string{"err when body parse"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// check email exist
	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Email == "" {
		errors := []string{"email or password incorrect"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		errors := []string{"email or password incorrect"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	//
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		log.Fatal(err)
		return
	}
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	var cart models.Cart
	database.DB.Where("user_id = ?", user.Id).First(&cart)
	response := map[string]string{
		"user_id":      strconv.Itoa(int(user.Id)),
		"name":         user.Name,
		"token":        token,
		"phone_number": user.PhoneNumber,
		"email":        user.Email,
		"sex":          user.Sex,
		"image":        user.Image,
		"birthday":     "",
		"cart_id":      "",
	}
	if user.Birthday != nil {
		response["birthday"] = user.Birthday.String()
	}
	if cart.ID != 0 {
		response["cart"] = strconv.Itoa(int(cart.ID))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

// admin

func RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	// parse body
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", data)
		return
	}

	// check password & confirm password
	if data["password"] != data["confirm"] {
		fmt.Fprintf(w, "password do not match")
		return
	}

	// check email existed
	var isExist models.Admin
	database.DB.Where("email = ?", data["email"]).First(&isExist)
	if isExist.Email != "" {
		fmt.Fprintf(w, "email exists")
		return
	}

	// hash password
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	var admin models.Admin
	admin.Name = data["name"]
	admin.Email = data["email"]
	admin.Password = password

	database.DB.Create(&admin)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Đăng ký tài khoản admin thành công <3")
}

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errors := []string{"err when body parse"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// check email exist
	var admin models.Admin

	database.DB.Where("email = ?", data["email"]).First(&admin)

	if admin.Email == "" {
		errors := []string{"email or password incorrect"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	if err := bcrypt.CompareHashAndPassword(admin.Password, []byte(data["password"])); err != nil {
		errors := []string{"email or password incorrect"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	//
	token, err := util.GenerateJwt(strconv.Itoa(int(admin.Id)))
	if err != nil {
		log.Fatal(err)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	response := map[string]string{
		"user_id": strconv.Itoa(int(admin.Id)),
		"name":    admin.Name,
		"token":   token,
		"admin":   "true",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	response := "logout success, see you again"
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}
