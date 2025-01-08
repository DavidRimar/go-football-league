package services

import (
	"api/internal/domain/interfaces"
	"api/internal/domain/models"
	"context"
)

type TeamService struct {
	repo interfaces.TeamRepository
}

func NewTeamService(repo interfaces.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) GetTeams(ctx context.Context) ([]models.Team, error) {
	return s.repo.GetAllTeams(ctx)
}
