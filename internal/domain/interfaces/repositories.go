package interfaces

import (
	"backend/internal/domain/models"
	"context"
)

type TeamRepository interface {
	GetAllTeams() ([]models.Team, error)
}

type FixturesRepository interface {
	GetAllFixtures(ctx context.Context) ([]models.Fixture, error)
	InsertFixtures(ctx context.Context, games []models.Fixture) error
}
