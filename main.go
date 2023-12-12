package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", mainHandler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
