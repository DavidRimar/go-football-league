package services

import (
	"backend/internal/domain/interfaces"
	"backend/internal/domain/models"
)

type TeamService struct {
	repo interfaces.TeamRepository
}

func NewTeamService(repo interfaces.TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) GetTeams() ([]models.Team, error) {
	return s.repo.GetAllTeams()
}
