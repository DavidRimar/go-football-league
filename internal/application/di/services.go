package di

import (
	"backend/internal/application/handlers"
	"backend/internal/application/services"
	"backend/internal/infrastructure/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	TeamHandler    *handlers.TeamHandler
	FixtureHandler *handlers.FixtureHandler
	SeedService    *services.DataSeederService
}

func InitializeServices(db *mongo.Database) *Container {
	// Initialize repositories
	teamRepo := repositories.NewTeamRepository(db)
	fixtureRepo := repositories.NewFixturesRepository(db)

	// Initialize services
	teamService := services.NewTeamService(teamRepo)
	fixtureService := services.NewFixtureService(fixtureRepo)
	seedService := services.NewDataSeederService(teamRepo, fixtureRepo)

	// Initialize handlers
	teamHandler := handlers.NewTeamHandler(teamService)
	fixtureHandler := handlers.NewFixtureHandler(fixtureService)

	return &Container{
		TeamHandler:    teamHandler,
		FixtureHandler: fixtureHandler,
		SeedService:    seedService,
	}
}
