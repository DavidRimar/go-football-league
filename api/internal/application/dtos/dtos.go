package dtos

import (
	"api/internal/domain/models"
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

type GetTeamStatisticsDTO struct {
	Team           string `json:"team"`
	GamesPlayed    int    `json:"gamesPlayed"`
	Points         int    `json:"points"`
	Wins           int    `json:"wins"`
	Draws          int    `json:"draws"`
	Losses         int    `json:"losses"`
	GoalsScored    int    `json:"goalsScored"`
	GoalsConceded  int    `json:"goalsConceded"`
	GoalDifference int    `json:"goalDifference"`
}
