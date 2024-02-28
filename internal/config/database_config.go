package config

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseConfig struct {
	PostgresOneDatabase *PostgresOneDatabase
}

type PostgresOneDatabase struct {
	Connection *sql.DB
}

func NewDatabaseConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		PostgresOneDatabase: NewPostgresOneDatabase(envConfig),
	}
	return databaseConfig
}

func NewPostgresOneDatabase(envConfig *EnvConfig) *PostgresOneDatabase {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		envConfig.PostgresOne.User,
		envConfig.PostgresOne.Password,
		envConfig.PostgresOne.Host,
		envConfig.PostgresOne.Port,
		envConfig.PostgresOne.Database,
	)
	connection, err := sql.Open(
		"pgx",
		url,
	)
	if err != nil {
		panic(err)
	}
	return &PostgresOneDatabase{
		Connection: connection,
	}
}
