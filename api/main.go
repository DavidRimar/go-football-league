package main

import (
	"api/config"
	"api/internal/application/di"
	"api/internal/application/router"
	"api/internal/application/utils"
	database "api/internal/infrastructure/persistence"
	"api/internal/middleware"
	"log"
	"net/http"
	"time"

	_ "api/docs" // Import the generated docs
)

// @title Football League API
// @version 1.0
// @description This is the Football League API documentation for the Football League service.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY
func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Define a context with a timeout
	ctx, cancel := utils.NewContextWithTimeout(10 * time.Second)
	defer cancel()

	// MongoDB
	mongoDB := database.InitializeMongoDB(ctx, cfg.DatabaseURI, cfg.DatabaseName)
	defer mongoDB.Client.Disconnect(ctx)

	// Connect to RabbitMQ
	rabbitMQChannel := utils.ConnectToRabbitMQ(cfg.RabbitMQConnectionString)
	defer utils.CloseRabbitMQ()

	// Define the queue name
	queueName := "team-stats-queue"

	// Initialize services and DI container
	container := di.InitializeServices(mongoDB.Database, rabbitMQChannel, queueName)

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
