package di

import (
	"backend/internal/application/handlers"
	"backend/internal/application/services"
	"backend/internal/infrastructure/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	TeamHandler      *handlers.TeamHandler
	FixtureHandler   *handlers.FixtureHandler
	TeamStatsHandler *handlers.TeamStatsHandler
	SeedService      *services.DataSeederService
}

func InitializeServices(db *mongo.Database) *Container {
	// Initialize repositories
	teamRepo := repositories.NewTeamRepository(db)
	fixtureRepo := repositories.NewFixturesRepository(db)
	teamStatsRepo := repositories.NewTeamStatisticsRepository(db)

	// Initialize services
	teamService := services.NewTeamService(teamRepo)
	fixtureService := services.NewFixtureService(fixtureRepo)
	teamStatService := services.NewTeamStatsService(teamStatsRepo)
	seedService := services.NewDataSeederService(teamRepo, fixtureRepo, teamStatsRepo)

	// Initialize handlers
	teamHandler := handlers.NewTeamHandler(teamService)
	fixtureHandler := handlers.NewFixtureHandler(fixtureService, teamStatService)
	teamStatsHandler := handlers.NewTeamStatsHandler(teamStatService)

	return &Container{
		TeamHandler:      teamHandler,
		FixtureHandler:   fixtureHandler,
		TeamStatsHandler: teamStatsHandler,
		SeedService:      seedService,
	}
}
