package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/pottery-api/routers"
)

// HandleServer handles the server request
func HandlerServer() {
	router := mux.NewRouter()

	users := router.PathPrefix("/api/v1/users").Subrouter()

	// Initialize the routes
	users.Path("").Methods(
		http.MethodPost).HandlerFunc(routers.CreateUserHandler)
	users.Path("").Methods(
		http.MethodGet).HandlerFunc(routers.GetUsersHandler)
	users.Path("/user").Methods(
		http.MethodGet).HandlerFunc(routers.GetUserByIdHandler)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server is running in port: " + PORT)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe("localhost:"+PORT, handler))
}
