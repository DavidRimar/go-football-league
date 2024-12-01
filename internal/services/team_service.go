package services

import "backend/internal/models"

type TeamRepository interface {
	GetAllTeams() ([]models.Team, error)
}

type TeamService struct {
	repo TeamRepository
}

func NewTeamService(repo TeamRepository) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) GetTeams() ([]models.Team, error) {
	return s.repo.GetAllTeams()
}
