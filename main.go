package main

import (
	"log"
	"multi-tenant-go-app/config"
	"multi-tenant-go-app/routes"
	"multi-tenant-go-app/utils"
	"net/http"
)

func main() {
	// Initialize the database connection
	config.InitDB()

	// Start a worker pool with 5 workers
	utils.StartWorkerPool(5)

	// Register routes
	r := routes.RegisterRoutes()

	// Start the server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
