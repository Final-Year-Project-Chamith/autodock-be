package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var RedClient *redis.Client

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{ 
		Addr: "localhost:6379",
		DB:   0,               
	})
	RedClient = client

	_, err := RedClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully.")
	return RedClient
}