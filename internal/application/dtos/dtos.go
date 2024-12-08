package dtos

// to go swere else
type UpdateFixtureDTO struct {
	HomeScore *int64  `json:"homeScore"`
	AwayScore *int64  `json:"awayScore"`
	Status    *string `json:"status"`
}
