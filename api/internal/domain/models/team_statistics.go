package models

import "time"

type TeamStatistics struct {
	TeamID         string    `bson:"teamId" json:"teamId"`
	Team           string    `bson:"team" json:"team"`
	GamesPlayed    int       `bson:"gamesPlayed" json:"gamesPlayed"`
	Wins           int       `bson:"wins" json:"wins"`
	Draws          int       `bson:"draws" json:"draws"`
	Losses         int       `bson:"losses" json:"losses"`
	GoalsScored    int       `bson:"goalsScored" json:"goalsScored"`
	GoalsConceded  int       `bson:"goalsConceded" json:"goalsConceded"`
	GoalDifference int       `bson:"goalDifference" json:"goalDifference"`
	Points         int       `bson:"points" json:"points"`
	LastUpdated    time.Time `bson:"lastUpdated" json:"lastUpdated"`
}
