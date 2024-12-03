package models

type FixtureStatus string

const (
	StatusPlayed   FixtureStatus = "Played"
	StatusLive     FixtureStatus = "Live"
	StatusUpcoming FixtureStatus = "Upcoming"
)
