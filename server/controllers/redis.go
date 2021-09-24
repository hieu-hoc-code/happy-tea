package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/database"
	"server/util"

	"github.com/go-redis/redis/v8"
)

func VerifyCache(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("user_id")
		if err != nil {
			response := map[string]string{
				"msg": "cannot get user_id",
			}
			json.NewEncoder(w).Encode(&response)
			return
		}
		userID := cookie.Value

		val, err := database.Cache.Get(database.Ctx, userID).Bytes()
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		data := util.ToJson(val)
		json.NewEncoder(w).Encode(&data)
	}
}

func ProductCache(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()
		products, err := database.Cache.Get(database.Ctx, url).Bytes()
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("cache")
		fmt.Fprint(w, string(products))
	}
}

func ClearProductCache(redis redis.Client, ctx context.Context) {
	redis.FlushAll(ctx)
}
