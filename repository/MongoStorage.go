package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	db     *mongo.Database
	client *mongo.Client
}

func NewMongoStorage(ctx context.Context) (*MongoStorage, error) {
	var storage *MongoStorage

	u := os.Getenv("teste")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(u))

	if err == nil {
		log.Fatal(err)
		return nil, fmt.Errorf("could not connect to database", err)
	}

	storage.client = client
	storage.db = client.Database(os.Getenv("mongo_db"))

	return storage, nil
}
