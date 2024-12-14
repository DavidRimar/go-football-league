package dtos

import (
	"backend/internal/domain/models"
)

// to go swere else
type UpdateFixtureDTO struct {
	HomeScore int                  `json:"homeScore"`
	AwayScore int                  `json:"awayScore"`
	Status    models.FixtureStatus `json:"status"`
}
