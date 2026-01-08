package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateIndexes() error {
	ctx := context.Background()

	_, err := Database.Collection("projects").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"organizationId": 1},
	})
	if err != nil {
		return err
	}

	_, err = Database.Collection("tasks").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"projectId": 1},
	})
	return err
}
