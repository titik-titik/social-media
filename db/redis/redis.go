package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	Client *redis.Client
}

func NewRedisManager() *RedisManager {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	redisHost := os.Getenv("GATEWAY_HOST")
	redisPort := os.Getenv("GATEWAY_REDIS_PORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	password := os.Getenv("GATEWAY_REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})

	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("error connecting to Redis: %v", err)
	}

	fmt.Printf("Connected to Redis at: %s\n", redisAddr)

	return &RedisManager{Client: client}
}

func (rm *RedisManager) Close() error {
	if err := rm.Client.Close(); err != nil {
		return fmt.Errorf("error closing Redis client: %w", err)
	}
	return nil
}

func (rm *RedisManager) InsertData(redisKey string, data []byte, expirationTime time.Time) error {
	ctx := context.Background()
	expirationDuration := time.Until(expirationTime)
	err := rm.Client.SetEx(ctx, redisKey, data, expirationDuration).Err()
	if err != nil {
		return fmt.Errorf("error setting data to cache: %w", err)
	}
	return nil
}
