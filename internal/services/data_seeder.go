package services

import (
	"context"
	"log"

	"backend/internal/models"
	"backend/internal/utils"
)

// interface to be moved?
type FixturesRepository interface {
	GetAllFixtures(ctx context.Context) ([]models.Fixture, error)
	InsertFixtures(ctx context.Context, games []models.Fixture) error
}

type DataSeederService struct {
	teamRepository     TeamRepository
	fixturesRepository FixturesRepository
}

func NewDataSeederService(teamRepo TeamRepository, fixturesRepo FixturesRepository) *DataSeederService {
	return &DataSeederService{
		teamRepository:     teamRepo,
		fixturesRepository: fixturesRepo,
	}
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
