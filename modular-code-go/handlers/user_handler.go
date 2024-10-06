package handlers

import (
	"encoding/json"
	"fmt"
	"modular-code-go/models"
	"modular-code-go/services"
	"net/http"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	UserService *services.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// AddUserHandler handles the HTTP request to add a user
func (uh *UserHandler) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	uh.UserService.AddUser(user)
	fmt.Fprintf(w, "User added successfully!")
}

// GetUsersHandler returns all users
func (uh *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := uh.UserService.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
