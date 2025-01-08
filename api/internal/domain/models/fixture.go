package models

import "time"

type Fixture struct {
	ID         string        `bson:"_id"`
	GameweekId int           `bson:"gameweekId"`
	HomeTeamId string        `bson:"homeTeamId"` // References Team.ID
	AwayTeamId string        `bson:"awayTeamId"` // References Team.ID
	HomeTeam   string        `bson:"homeTeam"`
	AwayTeam   string        `bson:"awayTeam"`
	Date       time.Time     `bson:"date"`
	Status     FixtureStatus `bson:"status"`
	HomeScore  int           `bson:"homeScore"`
	AwayScore  int           `bson:"awayScore"`
}
