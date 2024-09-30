package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes the database connection and runs migrations.
func InitDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Database connected successfully.")

	// Run database migrations
	runMigrations()
}

// runMigrations applies the database migrations.
func runMigrations() {
	// Get absolute path to the migrations folder
	absPath, err := filepath.Abs("./migrations")
	if err != nil {
		log.Fatal("Failed to get absolute path for migrations:", err)
	}

	m, err := migrate.New(
		"file://"+absPath, // Use absolute path
		os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to initialize migrate:", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to apply migrations:", err)
	} else if err == migrate.ErrNoChange {
		fmt.Println("No new migrations to apply.")
	} else {
		fmt.Println("Migrations applied successfully.")
	}
}
