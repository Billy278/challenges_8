package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDBMongo() *mongo.Database {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://billy:billy@localhost:27020")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	return client.Database("DB_Books")
}
