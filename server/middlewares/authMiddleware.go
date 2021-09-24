package middlewares

import (
	"encoding/json"
	"net/http"
	"server/util"
	"time"
)

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			response := map[string]string{
				"msg": "unauthenticated",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&response)
			return
		}

		if _, err := util.ParseJwt(cookie.Value); err != nil {
			response := map[string]string{
				"msg": "unauthenticated when parse cookie",
			}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&response)
			return
		}

		reCookie := http.Cookie{
			Name:     "jwt",
			Value:    cookie.Value,
			Expires:  time.Now().Add(time.Hour * 24),
			HttpOnly: true,
		}
		http.SetCookie(w, &reCookie)

		next.ServeHTTP(w, r)
	}
}
