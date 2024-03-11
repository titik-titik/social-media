package config

import (
	"os"
)

type AppEnv struct {
	Host string
	Port string
}

type CockroachdbEnv struct {
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

type LoggerEnv struct {
	Level int8
}

type EnvConfig struct {
	App         *AppEnv
	Cockroachdb *CockroachdbEnv
	RedisOne    *RedisOneEnv
	Logger      *LoggerEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host: os.Getenv("GATEWAY_HOST"),
			Port: os.Getenv("GATEWAY_APP_PORT"),
		},
		Cockroachdb: &CockroachdbEnv{
			Host:     os.Getenv("GATEWAY_HOST"),
			Port:     os.Getenv("GATEWAY_COCKROACHDB_SQL_PORT"),
			User:     os.Getenv("GATEWAY_COCKROACHDB_USER"),
			Password: os.Getenv("GATEWAY_COCKROACHDB_PASSWORD"),
			Database: os.Getenv("GATEWAY_COCKROACHDB_DATABASE"),
		},
		RedisOne: &RedisOneEnv{
			Host:     os.Getenv("GATEWAY_HOST"),
			Port:     os.Getenv("GATEWAY_REDIS_PORT"),
			Password: os.Getenv("GATEWAY_REDIS_PASSWORD"),
		},
		Logger: &LoggerEnv{
			Level: 6,
		},
	}
	return envConfig
}
