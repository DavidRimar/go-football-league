package dtos

import "backend/internal/domain/models"

// to go swere else
type UpdateFixtureDTO struct {
	HomeScore *int64                `json:"homeScore"`
	AwayScore *int64                `json:"awayScore"`
	Status    *models.FixtureStatus `json:"status"`
}
