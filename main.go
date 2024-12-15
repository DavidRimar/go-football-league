package main

import (
	"backend/config"
	"backend/internal/application/di"
	"backend/internal/application/router"
	"backend/internal/application/utils"
	database "backend/internal/infrastructure/persistence"
	"backend/internal/middleware"
	"log"
	"net/http"
	"time"

	_ "backend/docs" // Import the generated docs
)

// @title Football League API
// @version 1.0
// @description This is the Football League API documentation for the Football League service.
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Define a context with a timeout
	ctx, cancel := utils.NewContextWithTimeout(10 * time.Second)
	defer cancel()

	// MongoDB
	mongoDB := database.InitializeMongoDB(ctx, cfg.DatabaseURI, cfg.DatabaseName)
	defer mongoDB.Client.Disconnect(ctx)

	// Registering services
	container := di.InitializeServices(mongoDB.Database)

	// Seed teams and fixtures
	container.SeedService.SeedData(ctx)

	// Initialize the router
	mux := router.NewRouter(container.TeamHandler, container.FixtureHandler, container.TeamStatsHandler)

	// Add CORS
	mux.Use(middleware.CORS)

	// Start the server
	log.Printf("Server is running on port %s...", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
