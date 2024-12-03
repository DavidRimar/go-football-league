package utils

import (
	"fmt"
	"time"

	"backend/internal/models"
)

func GenerateFixtures(teams []models.Team) []models.Fixture {

	n := len(teams)
	rounds := n - 1
	var games []models.Fixture
	startDate := time.Date(2025, time.September, 6, 15, 0, 0, 0, time.UTC)

	for round := 0; round < rounds; round++ {
		gameweekId := round + 1
		gameweekDate := startDate.AddDate(0, 0, round*7) // Add one week for each gameweek

		for i := 0; i < n/2; i++ {
			home := teams[i]
			away := teams[n-1-i]

			// Skip matches against "Bye"
			if home.ID == "dummy" || away.ID == "dummy" {
				continue
			}

			// Create a game object
			game := models.Fixture{
				ID:         fmt.Sprintf("GW%d_FXT%d", gameweekId, i+1),
				HomeTeam:   home.ID,
				AwayTeam:   away.ID,
				Status:     models.StatusUpcoming,
				Date:       gameweekDate,
				GameweekId: gameweekId,
			}
			games = append(games, game)
		}
		// Rotate teams (except the first one)
		teams = append([]models.Team{teams[0]}, append(teams[n-1:], teams[1:n-1]...)...)
	}

	return games
}
