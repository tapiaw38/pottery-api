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

	// Initialize the routes
	router.HandleFunc("/api/v1/users/", routers.CreateUserHandler).Methods(http.MethodPost)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server is running in port: " + PORT)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe("localhost:"+PORT, handler))
}
