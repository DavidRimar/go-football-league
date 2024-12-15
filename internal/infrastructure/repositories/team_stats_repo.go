package repositories

import (
	"backend/internal/domain/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamStatisticsRepository struct {
	collection *mongo.Collection
}

func NewTeamStatisticsRepository(db *mongo.Database) *TeamStatisticsRepository {
	return &TeamStatisticsRepository{
		collection: db.Collection("teamStatistics"),
	}
}

func (r *TeamStatisticsRepository) GetAllTeamStatistics(ctx context.Context) ([]models.TeamStatistics, error) {

	var teamStats []models.TeamStatistics

	// Find all documents in the collection
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch team statistics: %w", err)
	}
	defer cursor.Close(ctx)

	// Iterate and decode each document into the slice
	for cursor.Next(ctx) {
		var stat models.TeamStatistics
		if err := cursor.Decode(&stat); err != nil {
			return nil, fmt.Errorf("failed to decode team statistics: %w", err)
		}
		teamStats = append(teamStats, stat)
	}

	return teamStats, nil
}

func (r *TeamStatisticsRepository) GetTeamStatistics(ctx context.Context, teamID string) (*models.TeamStatistics, error) {

	filter := bson.M{"teamId": teamID}

	var stats models.TeamStatistics
	err := r.collection.FindOne(ctx, filter).Decode(&stats)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			// set empty stats if not found
			stats = models.TeamStatistics{
				TeamID: teamID,
			}

			// return actual stats
			return &stats, nil
		}

		return nil, err
	}

	// return empty stats
	return &stats, nil
}

func (r *TeamStatisticsRepository) UpdateTeamStatistics(ctx context.Context, stats *models.TeamStatistics) error {

	filter := bson.M{"teamId": stats.TeamID}
	fieldsUpdate := bson.M{}

	fieldsUpdate["gamesPlayed"] = stats.GamesPlayed
	fieldsUpdate["wins"] = stats.Wins
	fieldsUpdate["draws"] = stats.Draws
	fieldsUpdate["losses"] = stats.Losses
	fieldsUpdate["goalsScored"] = stats.GoalsScored
	fieldsUpdate["goalsConceded"] = stats.GoalsConceded
	fieldsUpdate["goalDifference"] = stats.GoalDifference
	fieldsUpdate["points"] = stats.Points

	_, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": fieldsUpdate})
	if err != nil {
		return err
	}

	return nil
}

func (r *TeamStatisticsRepository) InsertTeamStatistics(ctx context.Context, stats []models.TeamStatistics) error {

	// Convert TeamStatistics to a slice of interface{} for MongoDB insertion
	records := make([]interface{}, len(stats))
	for i, stat := range stats {
		records[i] = stat
	}

	// Bulk insert
	_, err := r.collection.InsertMany(ctx, records)
	if err != nil {
		return fmt.Errorf("failed to insert team statistics: %w", err)
	}

	return nil
}
