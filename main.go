package main

import (
	"backend/config"
	"backend/internal/application/handlers"
	"backend/internal/application/router"
	"backend/internal/application/services"
	"backend/internal/infrastructure/repositories"
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

	// Registering services
	teamRepository := repositories.NewTeamRepository(db)
	fixtureRepository := repositories.NewFixturesRepository(db)

	teamService := services.NewTeamService(teamRepository)
	fixtureService := services.NewFixtureService(fixtureRepository)
	seedService := services.NewDataSeederService(teamRepository, fixtureRepository)

	teamHandler := handlers.NewTeamHandler(teamService)
	fixtureHandler := handlers.NewFixtureHandler(fixtureService)

	// Seed teams and fixtures
	if err := seedService.SeedTeams(ctx, "internal/data/teams.json"); err != nil {
		log.Fatalf("Teams seeding failed: %v", err)
	}

	if err := seedService.SeedFixtures(ctx); err != nil {
		log.Fatalf("Fixtures seeding failed: %v", err)
	}

	// Initialize the router
	mux := router.NewRouter(teamHandler, fixtureHandler)

	// Start the server
	log.Printf("Server is running on port %s...", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
