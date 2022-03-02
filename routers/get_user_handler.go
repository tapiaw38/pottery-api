package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tapiaw38/pottery-api/database"
)

// GetUsersHandler handles the request to get all users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	users, err := database.GetUsers(ctx)

	if err != nil {
		http.Error(w, "An error occurred when trying to get users "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)

}

// GetUserByIdHandler handles the request to get a user by id
func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	id := r.URL.Query().Get("id")

	user, err := database.GetUserById(ctx, id)

	if err != nil {
		http.Error(w, "An error occurred when trying to get user "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

// GetUserByUsernameHandler handles the request to get a user by username
func GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	username := r.URL.Query().Get("username")

	user, err := database.GetUserByUsername(ctx, username)

	if err != nil {
		http.Error(w, "An error occurred when trying to get user "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}
