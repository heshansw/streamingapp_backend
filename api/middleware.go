package api

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Add your authentication logic here

		// For example, check if a valid token is present in the request header

		// If authentication fails, respond with unauthorized status
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		//return

		// If authentication succeeds, call the next handler
		next.ServeHTTP(w, r)
	}
}
