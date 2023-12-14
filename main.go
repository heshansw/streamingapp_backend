package main

import (
	"backendapi/api"
	"backendapi/models"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := api.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDatabase()
	server.ListenAndServe()
}
