package repositories

import (
	"context"
	"hexagonal/internals/core/ports"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MongoClientTimeout = 5
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(conn string) *UserRepository {

	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		return nil
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil
	}
	return &UserRepository{
		client:     client,
		database:   client.Database("sample_airbnb"),
		collection: client.Database("sample_airbnb").Collection("listingsAndReviews"),
	}
}

// Login implements ports.UserRepository
func (r *UserRepository) Login(email string, password string) error {
	//Here your code for login in mongo database
	return nil
}

// Register implements ports.UserRepository
func (r *UserRepository) Register(email string, password string) error {
	//Here your code for save in mongo database
	return nil
}

var _ ports.UserRepository = (*UserRepository)(nil)
