package main

import (
	"log"
	"multi-tenant-go-app/config"
	"multi-tenant-go-app/routes"
	"net/http"
)

func main() {
	// Initialize database connection
	config.InitDB()

	// Register routes
	r := routes.RegisterRoutes()

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
