package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedClient *redis.Client

func NewRedisClient() *redis.Client {

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
		log.Println("REDIS_ADDR not set, defaulting to", redisAddr)
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr, // Redis server address
		DB:   0,         // Default database
	})

	RedClient = client

	_, err := RedClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis at %s: %v", redisAddr, err)
	}

	log.Printf("Connected to Redis successfully at %s.\n", redisAddr)
	return RedClient
}
