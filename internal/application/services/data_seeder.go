package services

import (
	"context"
	"log"

	"backend/internal/application/utils"
	"backend/internal/domain/interfaces"
)

type DataSeederService struct {
	teamRepository     interfaces.TeamRepository
	fixturesRepository interfaces.FixturesRepository
}

func NewDataSeederService(teamRepo interfaces.TeamRepository, fixturesRepo interfaces.FixturesRepository) *DataSeederService {
	return &DataSeederService{
		teamRepository:     teamRepo,
		fixturesRepository: fixturesRepo,
	}
}

func (s *DataSeederService) SeedTeams(ctx context.Context, filePath string) error {

	seedTeams := utils.LoadTeamsFromJSON(filePath)

	log.Printf("Teams to insert: %+v\n", seedTeams)

	teams, err := s.teamRepository.GetAllTeams()
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

func (s *DataSeederService) SeedFixtures(ctx context.Context) error {

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
	teams, err := s.teamRepository.GetAllTeams()
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
