package api

import (
	"backendapi/models"
	"backendapi/utils"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate
var jwtKey = []byte("secr")

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Token struct {
	Expires time.Time `json:"expires"`
	Token   string    `json:"token"`
}

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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cred Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)

	if err != nil {
		http.Error(w, "Error Occured", http.StatusBadRequest)
		return
	}

	var user models.Users
	result := models.DB.First(&user, "username = ?", cred.Username)

	if result.RowsAffected > 0 {
		chPassword := utils.CheckPassword(user.Password, cred.Password)

		if chPassword {
			expirationTime := time.Now().Add(5 * time.Minute)

			claims := &Claims{
				Username: cred.Username,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(expirationTime),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, errSig := token.SignedString(jwtKey)

			if errSig != nil {
				http.Error(w, "JWT Error", http.StatusBadRequest)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			var tokenData Token
			tokenData.Token = tokenString
			tokenData.Expires = expirationTime

			json.NewEncoder(w).Encode(tokenData)
			return
		} else {
			http.Error(w, "Password Error", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "Password Error", http.StatusBadRequest)
		return
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.UserInput

	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password, error := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	if error == nil {
		userNew := &models.Users{
			Username: input.Username,
			Password: string(password),
		}

		models.DB.Create(userNew)

		if userNew.UserId > 0 {
			json.NewEncoder(w).Encode(userNew)
		} else {
			http.Error(w, "Duplicate user name", http.StatusBadRequest)
		}

		return
	} else {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

}
