package config

import "os"

type AppEnv struct {
	Port string
}

type MariadbOneEnv struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

type RedisOneEnv struct {
	Host     string
	Port     string
	Password string
}

type EnvConfig struct {
	AppEnv        *AppEnv
	MariadbOneEnv *MariadbOneEnv
	RedisOneEnv   *RedisOneEnv
}

func NewEnvConfig() *EnvConfig {
	envConfig := &EnvConfig{
		AppEnv: &AppEnv{
			Port: os.Getenv("APP_PORT"),
		},
		MariadbOneEnv: &MariadbOneEnv{
			Host:     os.Getenv("MARIADB_ONE_HOST"),
			Port:     os.Getenv("MARIADB_ONE_PORT"),
			Database: os.Getenv("MARIADB_ONE_DATABASE"),
			User:     os.Getenv("MARIADB_ONE_USER"),
			Password: os.Getenv("MARIADB_ONE_PASSWORD"),
		},
		RedisOneEnv: &RedisOneEnv{
			Host:     os.Getenv("REDIS_ONE_HOST"),
			Port:     os.Getenv("REDIS_ONE_PORT"),
			Password: os.Getenv("REDIS_ONE_PASSWORD"),
		},
	}
	return envConfig
}
