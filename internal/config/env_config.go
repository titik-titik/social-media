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
			Host: os.Getenv("GATEWAY_HOST"),
			Port: os.Getenv("GATEWAY_APP_PORT"),
		},
		CockroachdbOne: &CockroachdbOneEnv{
			Host:     os.Getenv("GATEWAY_HOST"),
			Port:     os.Getenv("GATEWAY_COCKROACHDB_SQL_PORT"),
			User:     os.Getenv("GATEWAY_COCKROACHDB_USER"),
			Password: os.Getenv("GATEWAY_COCKROACHDB_PASSWORD"),
			Database: os.Getenv("GATEWAY_COCKROACHDB_DATABASE"),
		},
		RedisOne: &RedisOneEnv{
			Host:     os.Getenv("GATEWAY_HOST"),
			Port:     os.Getenv("REDIS_GATEWAY_PORT"),
			Password: os.Getenv("REDIS_GATEWAY_PASSWORD"),
		},
	}
	return envConfig
}
