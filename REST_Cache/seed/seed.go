package seed

import (
	"context"
	"log"
	"time"

	models "task-manager/collections"
	"task-manager/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedData() {
	ctx := context.Background()

	org := models.Organization{
		ID:        primitive.NewObjectID().Hex(),
		Name:      "Acme Corp",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	orgRes, err := db.Database.Collection("organizations").InsertOne(ctx, org)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted Organization:", orgRes.InsertedID)

	project := models.Project{
		ID:             primitive.NewObjectID().Hex(),
		Name:           "Internal Tooling",
		OrganizationID: org.ID,
		Status:         "in-progress",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	projRes, err := db.Database.Collection("projects").InsertOne(ctx, project)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted Project:", projRes.InsertedID)

	task := models.Task{
		ID:        primitive.NewObjectID().Hex(),
		Title:     "Design database schema",
		ProjectID: project.ID,
		Status:    "pending",
		Priority:  "high",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	taskRes, err := db.Database.Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted Task:", taskRes.InsertedID)
}
