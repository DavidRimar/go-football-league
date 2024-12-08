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

func (s *FixturesService) UpdateFixture(ctx context.Context, fixtureID string, fixture dtos.UpdateFixtureDTO) error {

	if fixtureID == "" {
		return errors.New("fixture ID cannot be empty")
	}

	// Fetch the existing fixture
	fixture, err := s.repo.GetFixtureByID(ctx, fixtureID)
	if err != nil {
		return err
	}

	fixture.HomeScore = *dto.HomeScore
	fixture.AwayScore = *dto.AwayScore
	fixture.Status = models.StatusPlayed // only allow to update Final Score

	return s.repo.UpdateFixture(ctx, fixtureID, fixture)
}
