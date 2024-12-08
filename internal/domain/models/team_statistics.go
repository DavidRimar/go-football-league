package models

import "time"

type TeamStatistics struct {
	TeamID         string    `bson:"teamId" json:"teamId"` // FK to Teams
	GamesPlayed    int       `bson:"numberOfGames" json:"numberOfGames"`
	Wins           int       `bson:"wins" json:"wins"`
	Draws          int       `bson:"draws" json:"draws"`
	Losses         int       `bson:"losses" json:"losses"`
	GoalsScored    int       `bson:"goalsScored" json:"goalsScored"`
	GoalsConceded  int       `bson:"goalsConceded" json:"goalsConceded"`
	GoalDifference int       `bson:"goalDifference" json:"goalDifference"`
	Points         int       `bson:"points" json:"points"`
	LastUpdated    time.Time `bson:"lastUpdated" json:"lastUpdated"`
}
