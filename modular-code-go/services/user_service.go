package services

import (
	"fmt"
	"modular-code-go/models"
)

// UserService handles the business logic for users
type UserService struct {
	Users []models.User
}

// NewUserService creates a new instance of UserService
func NewUserService() *UserService {
	return &UserService{Users: []models.User{}}
}

// AddUser adds a new user to the system
func (us *UserService) AddUser(user models.User) {
	us.Users = append(us.Users, user)
	fmt.Printf("Added new user: %s\n", user.Name)
}

// GetAllUsers returns all the users in the system
func (us *UserService) GetAllUsers() []models.User {
	return us.Users
}
