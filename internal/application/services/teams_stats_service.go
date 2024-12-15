package services

import (
	"backend/internal/application/dtos"
	"backend/internal/infrastructure/repositories"
	"context"
)

type TeamStatsService struct {
	teamStatsRepo *repositories.TeamStatisticsRepository
}

func NewTeamStatsService(teamStatsRepo *repositories.TeamStatisticsRepository) *TeamStatsService {
	return &TeamStatsService{teamStatsRepo: teamStatsRepo}
}

func (s *TeamStatsService) GetTeamStatistics(ctx context.Context) ([]dtos.GetTeamStatisticsDTO, error) {

	stats, err := s.teamStatsRepo.GetAllTeamStatistics(ctx)
	if err != nil {
		return nil, err
	}

	var dtoStats []dtos.GetTeamStatisticsDTO
	for _, stat := range stats {
		dtoStats = append(dtoStats, dtos.GetTeamStatisticsDTO{
			Team:           stat.Team,
			GamesPlayed:    stat.GamesPlayed,
			Wins:           stat.Wins,
			Draws:          stat.Draws,
			Losses:         stat.Losses,
			GoalsScored:    stat.GoalsScored,
			GoalsConceded:  stat.GoalsConceded,
			GoalDifference: stat.GoalDifference,
			Points:         stat.Points,
		})
	}

	return dtoStats, nil
}

func (s *TeamStatsService) UpdateTeamStatistics(ctx context.Context, fixture dtos.UpdateTeamStatsDTO) error {

	// Fetch statistics for both teams
	homeStats, err := s.teamStatsRepo.GetTeamStatistics(ctx, fixture.HomeTeamId)
	if err != nil {
		return err
	}

	awayStats, err := s.teamStatsRepo.GetTeamStatistics(ctx, fixture.AwayTeamId)
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
	homeStats.GoalDifference += (fixture.HomeScore - fixture.AwayScore)
	awayStats.GoalsScored += fixture.AwayScore
	awayStats.GoalsConceded += fixture.HomeScore
	awayStats.GoalDifference += (fixture.AwayScore - fixture.HomeScore)

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
