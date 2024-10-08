package main

import (
	"log"
	"net/http"
	"secure-file-upload/database"
	"secure-file-upload/handlers"
	"secure-file-upload/middleware"
	"secure-file-upload/utils"
)

func main() {
	// Initialize encryption (read the key from the environment)
	err := utils.InitCrypto()
	if err != nil {
		log.Fatalf("Failed to initialize encryption: %v", err)
	}

	// Connect to the PostgreSQL database
	database.ConnectDB()

	// Routes for file upload and download
	http.Handle("/upload", middleware.RateLimitMiddleware(http.HandlerFunc(handlers.UploadFile)))
	http.Handle("/download", middleware.RateLimitMiddleware(http.HandlerFunc(handlers.DownloadFile)))

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
