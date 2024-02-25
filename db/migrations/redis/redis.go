package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client // RedisClient is the instance of the Redis client
var RedisAddr string          // Variable to store the Redis address
var RedisKey string           // RedisKey is the key for storing user data in cache

// InitializeRedis sets up the connection to Redis
func InitializeRedis(envPath string) *redis.Client {
	if err := godotenv.Load(envPath); err != nil { // Using relative path
		log.Fatalf("Error loading .env file in caching: %v", err)
	}

	RedisAddr = os.Getenv("REDIS_ADDR")
	RedisPassword := os.Getenv("REDIS_PW")
	RedisDBStr := os.Getenv("REDIS_DB")

	RedisDB, err := strconv.Atoi(RedisDBStr)
	if err != nil {
		fmt.Println("Error converting Redis DB to integer:", err)
		os.Exit(1)
	}

	options := &redis.Options{
		Addr:     RedisAddr,
		Password: RedisPassword,
		DB:       RedisDB,
	}
	RedisClient = redis.NewClient(options)

	// Test connection and authentication to Redis
	ctx := context.Background()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		// Stop the program or handle the error according to your needs
		return nil
	}

	fmt.Printf("Connected to Redis at: %s\n", RedisAddr)
	return RedisClient
}

// CloseRedis closes the Redis connection
func CloseRedis() error {
	if err := RedisClient.Close(); err != nil {
		fmt.Println("Error closing Redis client:", err)
		return err
	}
	return nil
}

// SetCached stores data in cache with a specific Redis key and expiration time
func SetCached(redisKey string, data []byte, expirationTime time.Time) error {
	ctx := context.Background()
	expirationDuration := time.Until(expirationTime)

	err := RedisClient.SetEx(ctx, redisKey, data, expirationDuration).Err()
	if err != nil {
		return fmt.Errorf("error setting data to cache: %v", err)
	}

	return nil
}
