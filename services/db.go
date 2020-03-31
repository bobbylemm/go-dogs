package services

import (
	"context"
	"go-dogs/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Database
}

func ConnectToDB(ctx context.Context, config config.Config) (*DB, error) {
	clientOptions := options.Client().ApplyURI(config.DBUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	dogsDB := client.Database("dogs")

	return &DB{dogsDB}, nil
}


