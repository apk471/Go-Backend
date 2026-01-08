package models

import "time"

type Task struct {
	ID          string     `bson:"_id,omitempty" json:"id"`
	Title       string     `bson:"title" json:"title"`
	ProjectID   string     `bson:"projectId" json:"projectId"`
	Status      string     `bson:"status" json:"status"`
	Priority    string     `bson:"priority" json:"priority"`
	AssignedTo  *string    `bson:"assignedTo,omitempty" json:"assignedTo,omitempty"`
	AssignedAt  *time.Time `bson:"assignedAt,omitempty" json:"assignedAt,omitempty"`
	Description *string    `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time  `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time  `bson:"updatedAt" json:"updatedAt"`
}
