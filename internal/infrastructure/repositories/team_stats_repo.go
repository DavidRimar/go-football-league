package repositories

import (
	"backend/internal/domain/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamStatisticsRepository struct {
	collection *mongo.Collection
}

func NewTeamStatisticsRepository(collection *mongo.Collection) *TeamStatisticsRepository {
	return &TeamStatisticsRepository{collection: collection}
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
	update := bson.M{"$set": stats}

	_, err := r.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}
