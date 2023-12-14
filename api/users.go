package api

import (
	"backendapi/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var users []models.Users
	result := models.DB.Find(&users)

	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode(users)
	} else {
		json.NewEncoder(w).Encode(result.Error)
	}

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.Users

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if _ == nil {
		userNew := models.Users{
			Username: user.Username,
			Password: string(password),
		}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

}
