package middlewares

import (
	"net/http"

	"github.com/tapiaw38/pottery-api/database"
)

// CheckDB checks if the database connection is working
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := database.CheckConnection()

		if !err {
			http.Error(w, "An error occurred when trying to connect to the database", 500)
			return
		}

		next.ServeHTTP(w, r)
	})
}
