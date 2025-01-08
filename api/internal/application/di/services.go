package di

import (
	"api/internal/application/handlers"
	"api/internal/application/services"
	publisher "api/internal/infrastructure/events"
	"api/internal/infrastructure/repositories"

	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Container struct {
	TeamHandler      *handlers.TeamHandler
	FixtureHandler   *handlers.FixtureHandler
	TeamStatsHandler *handlers.TeamStatsHandler
	SeedService      *services.DataSeederService
}

func InitializeServices(db *mongo.Database, rabbitMQChannel *amqp091.Channel, queueName string) *Container {
	// Initialize repositories
	teamRepo := repositories.NewTeamRepository(db)
	fixtureRepo := repositories.NewFixturesRepository(db)
	teamStatsRepo := repositories.NewTeamStatisticsRepository(db)

	// Initialize services
	teamService := services.NewTeamService(teamRepo)
	fixtureService := services.NewFixtureService(fixtureRepo)
	teamStatService := services.NewTeamStatsService(teamStatsRepo)
	seedService := services.NewDataSeederService(teamRepo, fixtureRepo, teamStatsRepo)

	// Initialize the EventPublisher
	eventPublisher := publisher.NewRabbitMQPublisher(rabbitMQChannel, queueName)

	// Initialize handlers
	teamHandler := handlers.NewTeamHandler(teamService)
	fixtureHandler := handlers.NewFixtureHandler(fixtureService, teamStatService, eventPublisher)
	teamStatsHandler := handlers.NewTeamStatsHandler(teamStatService)

	return &Container{
		TeamHandler:      teamHandler,
		FixtureHandler:   fixtureHandler,
		TeamStatsHandler: teamStatsHandler,
		SeedService:      seedService,
	}
}
