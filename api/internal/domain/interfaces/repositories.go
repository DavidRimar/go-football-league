package interfaces

import (
	"api/internal/domain/models"
	"context"
)

type TeamRepository interface {
	GetAllTeams(ctx context.Context) ([]models.Team, error)
	InsertTeams(ctx context.Context, teams []models.Team) error
}

type FixturesRepository interface {
	GetAllFixtures(ctx context.Context) ([]models.Fixture, error)
	GetFixturesByGameweek(ctx context.Context, gameweekId int) ([]models.Fixture, error)
	InsertFixtures(ctx context.Context, fixtures []models.Fixture) error
	UpdateFixture(ctx context.Context, fixtureID string, fixture *models.Fixture) error
	GetFixtureByID(ctx context.Context, fixtureID string) (*models.Fixture, error)
}

type TeamStatsRepository interface {
	GetTeamStatistics(ctx context.Context, teamID string) (*models.TeamStatistics, error)
	UpdateTeamStatistics(ctx context.Context, stats *models.TeamStatistics) error
	InsertTeamStatistics(ctx context.Context, stats []models.TeamStatistics) error
	GetAllTeamStatistics(ctx context.Context) ([]models.TeamStatistics, error)
}
