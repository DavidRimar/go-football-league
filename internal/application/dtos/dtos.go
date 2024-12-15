package dtos

import (
	"backend/internal/domain/models"
)

type UpdateFixtureDTO struct {
	HomeScore int                  `json:"homeScore"`
	AwayScore int                  `json:"awayScore"`
	Status    models.FixtureStatus `json:"status"`
}

type UpdateTeamStatsDTO struct {
	HomeTeamId string `json:"homeTeamId"`
	AwayTeamId string `json:"awayTeamId"`
	HomeScore  int    `json:"homeScore"`
	AwayScore  int    `json:"awayScore"`
}
