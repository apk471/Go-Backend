package models

import "time"

type Project struct {
	ID             string    `bson:"_id,omitempty" json:"id"`
	Name           string    `bson:"name" json:"name"`
	OrganizationID string    `bson:"organizationId" json:"organizationId"`
	Status         string    `bson:"status" json:"status"`
	Description    *string   `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt      time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time `bson:"updatedAt" json:"updatedAt"`
}
