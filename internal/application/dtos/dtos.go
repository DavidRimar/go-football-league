package dtos

// to go swere else
type UpdateFixtureDTO struct {
	HomeScore int    `json:"homeScore"`
	AwayScore int    `json:"awayScore"`
	Status    string `json:"status"`
}
