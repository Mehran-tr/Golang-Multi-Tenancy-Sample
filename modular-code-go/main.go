package main

import (
	"modular-code-go/handlers"
	"modular-code-go/services"
	"modular-code-go/utils"
	"net/http"
)

func main() {
	// Create a UserService instance
	userService := services.NewUserService()

	// Create a UserHandler instance with dependency injection
	userHandler := handlers.NewUserHandler(userService)

	// Set up HTTP routes
	http.HandleFunc("/users", userHandler.GetUsersHandler)
	http.HandleFunc("/add-user", userHandler.AddUserHandler)

	// Start the server
	utils.LogInfo("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		utils.LogError(err)
	}
}
