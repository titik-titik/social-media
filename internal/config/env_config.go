package config

import "os"

type AppEnv struct {
	Host string
	Port string
}

type CockroachdbOneEnv struct {
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
	App            *AppEnv
	CockroachdbOne *CockroachdbOneEnv
	RedisOne       *RedisOneEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		App: &AppEnv{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		CockroachdbOne: &CockroachdbOneEnv{
			Host:     os.Getenv("COCKROACHDB_ONE_HOST"),
			Port:     os.Getenv("COCKROACHDB_ONE_SQL_PORT"),
			User:     os.Getenv("COCKROACHDB_ONE_USER"),
			Password: os.Getenv("COCKROACHDB_ONE_PASSWORD"),
			Database: os.Getenv("COCKROACHDB_ONE_DATABASE"),
		},
		RedisOne: &RedisOneEnv{
			Host:     os.Getenv("REDIS_ONE_HOST"),
			Port:     os.Getenv("REDIS_ONE_CLIENT_PORT"),
			Password: os.Getenv("REDIS_ONE_PASSWORD"),
		},
	}
	return envConfig
}
