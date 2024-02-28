package config

import "os"

type AppEnv struct {
	Host string
	Port string
}

type PostgresOneEnv struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type RedisOneEnv struct {
	Host     string
	Port     string
	Password string
}

type EnvConfig struct {
	App         *AppEnv
	PostgresOne *PostgresOneEnv
	RedisOne    *RedisOneEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		PostgresOne: &PostgresOneEnv{
			Host:     os.Getenv("POSTGRES_ONE_HOST"),
			Port:     os.Getenv("POSTGRES_ONE_PORT"),
			User:     os.Getenv("POSTGRES_ONE_USER"),
			Password: os.Getenv("POSTGRES_ONE_PASSWORD"),
			Database: os.Getenv("POSTGRES_ONE_DATABASE"),
		},
		RedisOne: &RedisOneEnv{
			Host:     os.Getenv("REDIS_ONE_HOST"),
			Port:     os.Getenv("REDIS_ONE_PORT"),
			Password: os.Getenv("REDIS_ONE_PASSWORD"),
		},
	}
	return envConfig
}
