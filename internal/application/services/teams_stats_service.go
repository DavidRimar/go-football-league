package services

import (
	"backend/internal/domain/models"
	"backend/internal/infrastructure/repositories"
	"context"
)

type TeamStatsService struct {
	teamStatsRepo *repositories.TeamStatisticsRepository
}

func NewTeamStatsService(teamStatsRepo *repositories.TeamStatisticsRepository) *TeamStatsService {
	return &TeamStatsService{teamStatsRepo: teamStatsRepo}
}

func (s *TeamStatsService) GetTeamStatistics(ctx context.Context) ([]models.TeamStatistics, error) {
	return s.teamStatsRepo.GetAllTeamStatistics(ctx)
}

func (s *TeamStatsService) UpdateTeamStatistics(ctx context.Context, fixture models.Fixture) error {

	// Fetch statistics for both teams
	homeStats, err := s.teamStatsRepo.GetTeamStatistics(ctx, fixture.HomeTeam)
	if err != nil {
		return err
	}

	awayStats, err := s.teamStatsRepo.GetTeamStatistics(ctx, fixture.AwayTeam)
	if err != nil {
		return err
	}

	// Increment stats
	homeStats.GamesPlayed++
	awayStats.GamesPlayed++

	if fixture.HomeScore > fixture.AwayScore {
		homeStats.Wins++
		awayStats.Losses++
		homeStats.Points += 3
	} else if fixture.HomeScore < fixture.AwayScore {
		awayStats.Wins++
		homeStats.Losses++
		awayStats.Points += 3
	} else {
		homeStats.Draws++
		awayStats.Draws++
		homeStats.Points++
		awayStats.Points++
	}

	homeStats.GoalsScored += fixture.HomeScore
	homeStats.GoalsConceded += fixture.AwayScore
	awayStats.GoalsScored += fixture.AwayScore
	awayStats.GoalsConceded += fixture.HomeScore

	// Save updated statistics
	err = s.teamStatsRepo.UpdateTeamStatistics(ctx, homeStats)
	if err != nil {
		return err
	}

	err = s.teamStatsRepo.UpdateTeamStatistics(ctx, awayStats)
	if err != nil {
		return err
	}

	return nil
}
