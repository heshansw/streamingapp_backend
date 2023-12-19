package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/users/register", RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", LoginUser).Methods("POST")

	privateRouter := router.PathPrefix("/").Subrouter()
	privateRouter.Use(AuthMiddleware)
	privateRouter.HandleFunc("/users", GetAllUsers).Methods("GET")

	privateRouter.HandleFunc("/upload", UploadVideo).Methods("POST")

	return router
}
