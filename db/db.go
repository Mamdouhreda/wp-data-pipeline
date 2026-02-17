package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DatabaseName   = "wpdata"
	CollectionName = "posts"
)

// ConnectToDB initializes and returns a MongoDB client
func ConnectToDB() (*mongo.Client, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		// Fallback to authenticated URI matching docker-compose defaults
		uri = "mongodb://admin:password@localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

