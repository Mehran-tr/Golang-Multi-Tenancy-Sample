package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB is the global database connection
var DB *sql.DB

// ConnectDB initializes the PostgreSQL connection using environment variables
func ConnectDB() {
	// Load .env file one level up
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the PostgreSQL config from the environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// PostgreSQL connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open connection
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the PostgreSQL database: %v", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the PostgreSQL database: %v", err)
	}

	// Create the files table if it doesn't exist
	createFilesTable := `
	CREATE TABLE IF NOT EXISTS files (
		id SERIAL PRIMARY KEY,
		original_name TEXT NOT NULL,
		encrypted_path TEXT NOT NULL
	);`
	if _, err := DB.Exec(createFilesTable); err != nil {
		log.Fatalf("Error creating files table: %v", err)
	}

	fmt.Println("Connected to the PostgreSQL database and files table is ready.")
}
