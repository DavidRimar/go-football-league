package main

import (
	"backend/config"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/services"
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Define a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DatabaseURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(cfg.DatabaseName)

	// Initialize layers
	repo := repositories.NewTeamRepository(db)
	service := services.NewTeamService(repo)
	handler := handlers.NewTeamHandler(service)

	// Define routes
	http.HandleFunc("/api/teams", handler.GetTeams)

	// Start the server
	log.Printf("Server is running on port %s...", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
