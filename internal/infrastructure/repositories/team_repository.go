package repositories

import (
	"backend/internal/domain/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamRepository struct {
	collection *mongo.Collection
}

func NewTeamRepository(db *mongo.Database) *TeamRepository {
	return &TeamRepository{
		collection: db.Collection("teams"),
	}
}

func (r *TeamRepository) GetAllTeams(ctx context.Context) ([]models.Team, error) {

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error querying teams collection: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var teams []models.Team
	if err = cursor.All(ctx, &teams); err != nil {
		log.Printf("Error decoding cursor: %v", err)
		return nil, err
	}

	return teams, nil
}

func (r *TeamRepository) InsertTeams(ctx context.Context, teams []models.Team) error {

	var teamDocuments []interface{}

	for _, team := range teams {
		teamDocument := bson.M{
			"name":            team.Name,
			"stadium":         team.Stadium,
			"stadiumCapacity": team.StadiumCapacity,
		}
		teamDocuments = append(teamDocuments, teamDocument)
	}

	_, err := r.collection.InsertMany(ctx, teamDocuments)

	return err
}
