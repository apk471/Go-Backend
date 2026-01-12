package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Global Redis client
var Client *redis.Client

// Connect initializes the Redis connection
func Connect() error {
	// Create a new Redis client
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Default Redis DB
	})

	// Ping Redis to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return Client.Ping(ctx).Err()
}