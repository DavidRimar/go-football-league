package repositories

import (
	"backend/internal/domain/models"
	"context"
	"log"
	"time"

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

func (r *FixturesRepository) GetFixturesByGameweek(gameweekId int) ([]models.Fixture, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"gameweekId": gameweekId}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Printf("Error querying fixtures collection: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var fixtures []models.Fixture
	if err := cursor.All(ctx, &fixtures); err != nil {
		log.Printf("Error decoding fixtures cursor: %v", err)
		return nil, err
	}

	return fixtures, nil
}

func (r *FixturesRepository) InsertFixtures(ctx context.Context, games []models.Fixture) error {
	var gameInterfaces []interface{}
	for _, game := range games {
		gameInterfaces = append(gameInterfaces, game)
	}

	_, err := r.collection.InsertMany(ctx, gameInterfaces)
	return err
}
