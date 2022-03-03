package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tapiaw38/pottery-api/middlewares"
	user "github.com/tapiaw38/pottery-api/routers/user"
)

// HandleServer handles the server request
func HandlerServer() {
	router := mux.NewRouter()

	// Initialize the routes
	users := router.PathPrefix("/api/v1/users").Subrouter()

	// Routes for users
	users.Path("/register").Methods(
		http.MethodPost).HandlerFunc(middlewares.CheckDB(user.CreateUserHandler))
	users.Path("getall").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(user.GetUsersHandler))
	users.Path("/user").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(user.GetUserByIdHandler))
	users.Path("/profile").Methods(
		http.MethodGet).HandlerFunc(middlewares.CheckDB(user.GetUserByUsernameHandler))
	users.Path("/update").Methods(
		http.MethodPut).HandlerFunc(middlewares.CheckDB(user.UpdateUserHandler))
	users.Path("/delete").Methods(
		http.MethodDelete).HandlerFunc(middlewares.CheckDB(user.DeleteUserHandler))

	// Initialize the server
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	log.Println("Server is running in port: " + PORT)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe("localhost:"+PORT, handler))
}
