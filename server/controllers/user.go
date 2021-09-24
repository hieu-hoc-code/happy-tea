package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/database"
	"server/models"
	"server/util"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func User(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		response := map[string]string{
			"msg": "cannot get user_id",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}
	userID := cookie.Value
	var user models.User
	database.DB.Where("id = ?", userID).First(&user)
	if user.Id == 0 {
		json.NewEncoder(w).Encode("không có user id")
		response := map[string]string{
			"msg": "cannot found user_id",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	cookie, _ = r.Cookie("jwt")
	body := models.Cache{
		UserID:      strconv.Itoa(int(user.Id)),
		Name:        user.Name,
		Token:       cookie.Value,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Sex:         user.Sex,
		Birthday:    "",
		Image:       user.Image,
		CartID:      "",
	}
	if user.Birthday != nil {
		body.Birthday = user.Birthday.String()
	}

	var cart models.Cart
	database.DB.Where("user_id = ?", user.Id).First(&cart)
	if cart.ID != 0 {
		body.CartID = strconv.Itoa(int(cart.ID))
	}

	data, _ := json.Marshal(&body)

	cacheErr := database.Cache.Set(database.Ctx, userID, data, 10*time.Second).Err()
	if cacheErr != nil {
		fmt.Fprintf(w, "%v", cacheErr)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&body)
}

func UploadImageUser(w http.ResponseWriter, r *http.Request) {
	logrus.Info("upload img user")
	userID := r.FormValue("user_id")

	var data models.User

	database.DB.Where("id = ?", userID).First(&data)
	if data.Id == 0 {
		fmt.Fprintf(w, "khong tim thay user")
		return
	}

	// delete old image
	if data.Image != "" {
		filePath := data.Image[21:]
		filePath = "." + filePath
		util.RemoveFile(filePath)
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	tempFile, err := ioutil.TempFile("static", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	url := "http://localhost:3000/" + tempFile.Name()

	database.DB.Model(data).Update("image", url)

	response := map[string]string{
		"url": url,
	}
	//clear cache
	ClearProductCache(*database.Cache, database.Ctx)
	json.NewEncoder(w).Encode(&response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// get id
	cookie, err := r.Cookie("jwt")
	if err != nil {
		response := map[string]string{
			"msg": "unauthenticated",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}
	parseCookie := cookie.Value

	valueCookie, err := util.ParseJwt(parseCookie)
	if err != nil {
		response := map[string]string{
			"msg": "unauthenticated",
		}
		json.NewEncoder(w).Encode(&response)
		return
	}

	// parse
	type User struct {
		Name            string     `json:"name"`
		PhoneNumber     string     `json:"phone_number"`
		Sex             string     `json:"sex"`
		Birthday        *time.Time `json:"birthday"`
		Password        string     `json:"password"`
		NewPassword     string     `json:"new_password"`
		ConfirmPassword string     `json:"confirm_password"`
	}

	var data User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprintf(w, "error when body parse %v", err)
		return
	}

	// check name
	if data.Name == "" {
		errors := []string{"cannot update"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// get user
	var user models.User
	database.DB.Where("id = ?", valueCookie).First(&user)
	if user.Email == "" {
		errors := []string{"cannot find account"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// update common infor
	user.Name = data.Name
	user.PhoneNumber = data.PhoneNumber
	if data.Sex == "male" || data.Sex == "female" {
		user.Sex = data.Sex
	}
	user.Birthday = data.Birthday

	if data.Password == "" {
		database.DB.Model(&user).Updates(user)

		response := "update completed"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&response)
		return
	}

	// check password & confirm password
	if data.Password != "" && data.ConfirmPassword != data.NewPassword {
		fmt.Fprintf(w, "password and confirm password incorrect")
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		errors := []string{"password incorrect"}
		response := map[string][]string{
			"errors": errors,
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 14)

	user.Password = password
	database.DB.Model(&user).Updates(user)

	// clear cookie
	clearCookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &clearCookie)

	msg := []string{"update completed", "bạn vừa thay đổi mật khẩu, vui lòng đăng nhập lại"}
	clearToken := []string{"true"}
	response := map[string][]string{
		"msg":         msg,
		"clear_token": clearToken,
	}
	//clear cache
	ClearProductCache(*database.Cache, database.Ctx)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&response)
}

func GetAdmin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user_id")
	if err != nil {
		response := map[string]string{
			"msg": "cannot get admin_id",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}
	adminID := cookie.Value
	var admin models.Admin
	database.DB.Where("id = ?", adminID).First(&admin)
	if admin.Id == 0 {
		json.NewEncoder(w).Encode("không có user id")
		response := map[string]string{
			"msg": "cannot found user_id",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&response)
		return
	}
	cookie, _ = r.Cookie("jwt")

	type Data struct {
		UserID string `json:"user_id"`
		Name   string `json:"name"`
		Token  string `json:"token"`
		Email  string `json:"email"`
		Admin  string `json:"admin"`
	}

	data := Data{
		UserID: adminID,
		Name:   admin.Name,
		Token:  cookie.Value,
		Email:  admin.Email,
		Admin:  "true",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&data)
}
