package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tokenString))
}

func validateCredentials(username, password string) bool {
	// ... (unchanged)

	return err == nil
}
