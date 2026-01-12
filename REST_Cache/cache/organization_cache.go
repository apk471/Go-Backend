package cache

import (
	"context"
	"encoding/json"
	"time"

	models "task-manager/collections"
)

// Cache TTL (5 minutes)
const organizationTTL = 5 * time.Minute

// GetOrganization retrieves organization from Redis
func GetOrganization(ctx context.Context, id string) (*models.Organization, error) {
	// Generate cache key
	key := "organization:" + id

	// Try fetching cached value
	val, err := Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	// Convert JSON back to struct
	var org models.Organization
	if err := json.Unmarshal([]byte(val), &org); err != nil {
		return nil, err
	}

	return &org, nil
}

// SetOrganization stores organization in Redis
func SetOrganization(ctx context.Context, org models.Organization) error {
	// Generate cache key
	key := "organization:" + org.ID

	// Convert struct to JSON
	data, err := json.Marshal(org)
	if err != nil {
		return err
	}

	// Store in Redis with TTL
	return Client.Set(ctx, key, data, organizationTTL).Err()
}

// DeleteOrganization removes organization from cache
func DeleteOrganization(ctx context.Context, id string) error {
	key := "organization:" + id
	return Client.Del(ctx, key).Err()
}