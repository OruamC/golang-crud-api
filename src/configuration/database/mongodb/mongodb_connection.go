package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"os"
)

var (
	MONGODB_URL      = "MONGODB_URL"
	MONGODB_USERS_DB = "MONGODB_USERS_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodbUri := os.Getenv(MONGODB_URL)
	mongodbUersdb := os.Getenv(MONGODB_USERS_DB)

	client, err := mongo.Connect(options.Client().ApplyURI(mongodbUri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodbUersdb), nil
}
