package repositories

import (
	"context"

	models "task-manager/collections"
	"task-manager/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetOrganizationByID(ctx context.Context, id string) (*models.Organization, error) {
	var org models.Organization

	err := db.Database.
		Collection("organizations").
		FindOne(ctx, bson.M{"_id": id}).
		Decode(&org)

	if err != nil {
		return nil, err
	}

	return &org, nil
}

func ListOrganizations(
	ctx context.Context,
	page int64,
	limit int64,
	status *string,
) ([]models.Organization, int64, error) {

	filter := bson.M{}
	if status != nil {
		filter["status"] = *status
	}

	opts := options.Find().
		SetSkip((page - 1) * limit).
		SetLimit(limit).
		SetSort(bson.M{"createdAt": -1})

	cursor, err := db.Database.
		Collection("organizations").
		Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var orgs []models.Organization
	if err := cursor.All(ctx, &orgs); err != nil {
		return nil, 0, err
	}

	total, err := db.Database.
		Collection("organizations").
		CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return orgs, total, nil
}

func CreateOrganization(ctx context.Context, org models.Organization) error {
	_, err := db.Database.
		Collection("organizations").
		InsertOne(ctx, org)
	return err
}

func UpdateOrganization(
	ctx context.Context,
	id string,
	update bson.M,
) error {
	res, err := db.Database.
		Collection("organizations").
		UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})

	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func DeleteOrganization(ctx context.Context, id string) error {
	res, err := db.Database.
		Collection("organizations").
		DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
