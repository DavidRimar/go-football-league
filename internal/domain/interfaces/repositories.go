package interfaces

import (
	"backend/internal/domain/models"
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
}
