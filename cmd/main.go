package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/router"
	"backend/internal/services"
	"context"
	"log"
	"net/http"
	"time"

	_ "backend/docs" // Import the generated docs

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title Football League API
// @version 1.0
// @description This is the Football League API documentation for the Football League service.
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Define a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.DatabaseURI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(cfg.DatabaseName)

	// Initialize layers
	repo := repositories.NewTeamRepository(db)
	service := services.NewTeamService(repo)
	handler := handlers.NewTeamHandler(service)

	// Initialize the router
	mux := router.NewRouter(handler)

	// Start the server
	log.Printf("Server is running on port %s...", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
