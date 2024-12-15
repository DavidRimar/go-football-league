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

type GetTeamStatisticsDTO struct {
	Team           string `json:"team"`
	GamesPlayed    int    `bson:"gamesPlayed" json:"gamesPlayed"`
	Points         int    `bson:"points" json:"points"`
	Wins           int    `bson:"wins" json:"wins"`
	Draws          int    `bson:"draws" json:"draws"`
	Losses         int    `bson:"losses" json:"losses"`
	GoalsScored    int    `bson:"goalsScored" json:"goalsScored"`
	GoalsConceded  int    `bson:"goalsConceded" json:"goalsConceded"`
	GoalDifference int    `bson:"goalDifference" json:"goalDifference"`
}
