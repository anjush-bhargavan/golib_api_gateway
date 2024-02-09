package services

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient represents a Redis client.
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient initializes a new Redis client.
func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server to check the connection
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &RedisClient{client: rdb}
}

// SetValue sets a key-value pair in Redis with an optional expiration time.
func (rc *RedisClient) SetValue(key, value string, expiration time.Duration) error {
	return rc.client.Set(context.Background(), key, value, expiration).Err()
}

// GetValue retrieves the value for a given key from Redis.
func (rc *RedisClient) GetValue(key string) (string, error) {
	return rc.client.Get(context.Background(), key).Result()
}
