package services

import (
	"backend/internal/domain/models"
	"backend/internal/infrastructure/repositories"
	"context"
)

type FixturesService struct {
	repo *repositories.FixturesRepository
}

func NewFixtureService(repo *repositories.FixturesRepository) *FixturesService {
	return &FixturesService{repo: repo}
}

func (s *FixturesService) GetFixturesByGameweek(ctx context.Context, gameweekId int) ([]models.Fixture, error) {
	return s.repo.GetFixturesByGameweek(ctx, gameweekId)
}
