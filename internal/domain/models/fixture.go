package models

import "time"

type Fixture struct {
	ID         string        `bson:"_id"`
	GameweekId int           `bson:"gameweekId"`
	HomeTeam   string        `bson:"homeTeam"` // References Team.ID
	AwayTeam   string        `bson:"awayTeam"` // References Team.ID
	Date       time.Time     `bson:"date"`
	Status     FixtureStatus `bson:"status"`
}
