package services

import (
	"backend/internal/application/dtos"
	"backend/internal/domain/models"
	"backend/internal/infrastructure/repositories"
	"context"
	"errors"
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

func (s *FixturesService) UpdateFixture(ctx context.Context, fixtureID string, dto dtos.UpdateFixtureDTO) error {

	if fixtureID == "" {
		return errors.New("fixture ID cannot be empty")
	}

	updatedFixture := models.Fixture{}
	updatedFixture.HomeScore = *dto.HomeScore
	updatedFixture.AwayScore = *dto.AwayScore
	updatedFixture.Status = models.StatusPlayed // only allow to update Final Score

	return s.repo.UpdateFixture(ctx, fixtureID, updatedFixture)
}
