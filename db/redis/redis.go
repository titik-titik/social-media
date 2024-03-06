package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	Client *redis.Client
}

func NewRedisConnection() (*RedisManager, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	redisHost := os.Getenv("REDIS_ONE_HOST")
	redisPort := os.Getenv("REDIS_ONE_PORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	password := os.Getenv("REDIS_ONE_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("error connecting to Redis: %w", err)
	}

	fmt.Printf("Connected to Redis at: %s\n", redisAddr)

	return &RedisManager{Client: client}, nil
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
