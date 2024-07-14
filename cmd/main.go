package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/wisnuuakbr/booking-to-go-golang/config"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/di"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/infrastructure/datastore"
	"github.com/wisnuuakbr/booking-to-go-golang/internal/usecases"

	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file located in the root directory
	envPath, err := filepath.Abs(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatalf("Error constructing .env file path: %v", err)
	}
	
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize configuration
	cfg := config.New()

	// Initialize database connection
	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	// Initialize repositories
	customerRepo := datastore.NewCustomerRepository(db)
	familyListRepo := datastore.NewFamilyListRepository(db)
	nationalityRepo := datastore.NewNationalityRepository(db)

	// Initialize use case
	customerUC := usecases.NewCustomerUseCase(customerRepo, familyListRepo, nationalityRepo)

	// Setup router
	router := di.SetupRouter(customerUC)

	// Start server
	log.Printf("Starting server on port %d...", cfg.App.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.Port), router))
}