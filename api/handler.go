package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	//router.HandleFunc("/users", GetAllUsers).Methods("GET")

	router.HandleFunc("/users/register", RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", LoginUser).Methods("POST")

	privateRouter := router.PathPrefix("/").Subrouter()
	privateRouter.Use(AuthMiddleware)
	privateRouter.HandleFunc("/users", GetAllUsers).Methods("GET")
	//router.HandleFunc("/quest/{id}",GetQuest).Methods("GET")
	//router.HandleFunc("/quest", CreateQuest).Methods("POST")
	//router.HandleFunc("/quest/{id}", UpdateQuest).Methods("PUT")
	//router.HandleFunc("/quest/{id}", DeleteQuest).Methods("DELETE")

	return router
}
