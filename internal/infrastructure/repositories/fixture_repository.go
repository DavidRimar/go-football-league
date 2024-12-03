package repositories

import (
	"backend/internal/domain/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FixturesRepository struct {
	collection *mongo.Collection
}

func NewFixturesRepository(db *mongo.Database) *FixturesRepository {
	return &FixturesRepository{
		collection: db.Collection("fixtures"),
	}
}

// GetAllFixtures fetches all games from the collection
func (r *FixturesRepository) GetAllFixtures(ctx context.Context) ([]models.Fixture, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var games []models.Fixture
	if err := cursor.All(ctx, &games); err != nil {
		return nil, err
	}

	return games, nil
}

// InsertFixtures inserts multiple games into the collection
func (r *FixturesRepository) InsertFixtures(ctx context.Context, games []models.Fixture) error {
	var gameInterfaces []interface{}
	for _, game := range games {
		gameInterfaces = append(gameInterfaces, game)
	}

	_, err := r.collection.InsertMany(ctx, gameInterfaces)
	return err
}
