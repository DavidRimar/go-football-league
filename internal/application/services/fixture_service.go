package services

import (
	"backend/internal/domain/models"
	"backend/internal/infrastructure/repositories"
)

type FixturesService struct {
	repo *repositories.FixturesRepository
}

func NewFixtureService(repo *repositories.FixturesRepository) *FixturesService {
	return &FixturesService{repo: repo}
}

func (s *FixturesService) GetFixturesByGameweek(gameweekId int) ([]models.Fixture, error) {
	return s.repo.GetFixturesByGameweek(gameweekId)
}
