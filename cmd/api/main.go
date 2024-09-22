package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/srhtak/go-ddd-api/internal/application"
	"github.com/srhtak/go-ddd-api/internal/infrastructure/persistence"
	handlers "github.com/srhtak/go-ddd-api/internal/interfaces/http"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Construct database connection string from environment variables
	dbConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Initialize repository
	repo, err := persistence.NewPostgresUserRepository(dbConnectionString)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// Initialize service
	userService := application.NewUserService(repo)

	// Initialize handler
	userHandler := handlers.NewUserHandler(userService)

	// Set up routes
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/auth", userHandler.AuthenticateUser).Methods("POST")

	// Start server
	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}