package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedClient *redis.Client

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379", // Redis server address
		DB:   0,                // Default DB (can be changed)
	})
	RedClient = client

	// Check Redis connection
	_, err := RedClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully.")
	return RedClient
}
