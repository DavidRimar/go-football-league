package utils

import (
	"backend/internal/domain/models"
	"fmt"
	"time"
)

func GenerateFixtures(teams []models.Team) []models.Fixture {

	n := len(teams)
	rounds := n - 1
	var fixtures []models.Fixture
	startDate := time.Date(2025, time.September, 6, 15, 0, 0, 0, time.UTC)

	// first half of fixtures
	for round := 0; round < rounds; round++ {
		gameweekId := round + 1
		gameweekDate := startDate.AddDate(0, 0, round*7) // Add one week for each gameweek

		for i := 0; i < n/2; i++ {
			home := teams[i]
			away := teams[n-1-i]

			homeFixture := models.Fixture{
				ID:         fmt.Sprintf("GW%d_FXT%d", gameweekId, i+1),
				HomeTeamId: home.ID,
				AwayTeamId: away.ID,
				Status:     models.StatusUpcoming,
				Date:       gameweekDate,
				GameweekId: gameweekId,
			}
			fixtures = append(fixtures, homeFixture)
		}

		// Rotate teams (except the first one)
		teams = append([]models.Team{teams[0]}, append(teams[n-1:], teams[1:n-1]...)...)
	}

	// Reverse fixtures
	reverseStartDate := startDate.AddDate(0, 0, (rounds*7)+6) // 6 weeks winter break

	for round := 0; round < rounds; round++ {
		gameweekId := rounds + round + 1
		gameweekDate := reverseStartDate.AddDate(0, 0, round*7)

		for i := 0; i < n/2; i++ {
			home := teams[n-1-i]
			away := teams[i]

			awayFixture := models.Fixture{
				ID:         fmt.Sprintf("gameweek%d_game%d", gameweekId, i+1),
				HomeTeamId: home.ID,
				AwayTeamId: away.ID,
				Status:     models.StatusUpcoming,
				Date:       gameweekDate,
				GameweekId: gameweekId,
			}
			fixtures = append(fixtures, awayFixture)
		}

		// Rotate teams (except the first one)
		teams = append([]models.Team{teams[0]}, append(teams[n-1:], teams[1:n-1]...)...)
	}

	return fixtures
}
