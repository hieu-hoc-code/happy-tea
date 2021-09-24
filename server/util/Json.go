package util

import (
	"encoding/json"
	"server/models"
)

func ToJson(val []byte) models.Cache {
	user := models.Cache{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}
	return user
}
