package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"backend/internal/application/utils"
	"backend/internal/domain/interfaces"
	"backend/internal/domain/models"
)

type DataSeederService struct {
	teamRepository      interfaces.TeamRepository
	fixturesRepository  interfaces.FixturesRepository
	teamStatsRepository interfaces.TeamStatsRepository
}

func NewDataSeederService(
	teamRepo interfaces.TeamRepository,
	fixturesRepo interfaces.FixturesRepository,
	teamStatsRepository interfaces.TeamStatsRepository,
) *DataSeederService {
	return &DataSeederService{
		teamRepository:      teamRepo,
		fixturesRepository:  fixturesRepo,
		teamStatsRepository: teamStatsRepository,
	}
}

func (s *DataSeederService) SeedData(ctx context.Context) error {
	if err := s.seedTeams(ctx, "internal/data/teams.json"); err != nil {
		log.Fatalf("Teams seeding failed: %v", err)
		return err
	}

	if err := s.seedFixtures(ctx); err != nil {
		log.Fatalf("Fixtures seeding failed: %v", err)
		return err
	}

	if err := s.seedTeamStats(ctx); err != nil {
		log.Fatalf("Team stats seeding failed: %v", err)
		return err
	}

	return nil
}

func (s *DataSeederService) seedTeams(ctx context.Context, filePath string) error {

	seedTeams := utils.LoadTeamsFromJSON(filePath)

	teams, err := s.teamRepository.GetAllTeams(ctx)
	if err != nil {
		return err
	}

	if len(teams) > 0 {
		log.Println("Teams already seeded.")
		return nil
	}

	err = s.teamRepository.InsertTeams(ctx, seedTeams)
	if err != nil {
		return err
	}

	log.Println("Teams seeded successfully!")

	return nil
}

func (s *DataSeederService) seedFixtures(ctx context.Context) error {

	// Check if fixtures already exist
	existingGames, err := s.fixturesRepository.GetAllFixtures(ctx)
	if err != nil {
		return err
	}

	if len(existingGames) > 0 {
		log.Println("Fixtures already exist in the database. Skipping generation.")
		return nil
	}

	// Fetch all teams
	teams, err := s.teamRepository.GetAllTeams(ctx)
	if err != nil {
		return err
	}

	// Generate fixtures
	games := utils.GenerateFixtures(teams)

	// Save games to the database
	err = s.fixturesRepository.InsertFixtures(ctx, games)
	if err != nil {
		return err
	}

	log.Println("Fixtures generated and inserted successfully!")
	return nil
}

func (s *DataSeederService) seedTeamStats(ctx context.Context) error {

	existingTeamStats, err := s.teamStatsRepository.GetAllTeamStatistics(ctx)
	if err != nil {
		log.Fatalf("Error fetching team statistics: %v", err)
	}

	if len(existingTeamStats) > 0 {
		fmt.Println("Team statistics already seeded")
		return nil
	}

	teams, err := s.teamRepository.GetAllTeams(ctx)
	if err != nil {
		return err
	}

	teamStats := createTeamStatisticsSeedData(teams)

	err = s.teamStatsRepository.InsertTeamStatistics(ctx, teamStats)
	if err != nil {
		return err
	}

	log.Println("Team stats initialised successfully!")
	return nil
}

func createTeamStatisticsSeedData(teams []models.Team) []models.TeamStatistics {
	seedTeamStats := make([]models.TeamStatistics, len(teams))

	for i, team := range teams {
		seedTeamStats[i] = models.TeamStatistics{
			TeamID:         team.ID,
			Team:           team.Name,
			GamesPlayed:    0,
			Wins:           0,
			Draws:          0,
			Losses:         0,
			GoalsScored:    0,
			GoalsConceded:  0,
			GoalDifference: 0,
			Points:         0,
			LastUpdated:    time.Now(),
		}
	}

	return seedTeamStats
}
