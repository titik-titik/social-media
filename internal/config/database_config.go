package config

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseConfig struct {
	CockroachdbOneDatabase *CockroachdbOneDatabase
}

type CockroachdbOneDatabase struct {
	Connection *sql.DB
}

func NewDatabaseConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		CockroachdbOneDatabase: NewCockroachdbOneDatabase(envConfig),
	}
	return databaseConfig
}

func NewCockroachdbOneDatabase(envConfig *EnvConfig) *CockroachdbOneDatabase {
	var url string
	if envConfig.CockroachdbOne.Password == "" {
		url = fmt.Sprintf(
			"postgresql://%s@%s:%s/%s",
			envConfig.CockroachdbOne.User,
			envConfig.CockroachdbOne.Host,
			envConfig.CockroachdbOne.Port,
			envConfig.CockroachdbOne.Database,
		)
	} else {
		url = fmt.Sprintf(
			"postgresql://%s@%s:%s/%s",
			envConfig.CockroachdbOne.User,
			envConfig.CockroachdbOne.Host,
			envConfig.CockroachdbOne.Port,
			envConfig.CockroachdbOne.Database,
		)
	}

	connection, err := sql.Open(
		"pgx",
		url,
	)
	if err != nil {
		panic(err)
	}

	cockroachdbOneDatabase := &CockroachdbOneDatabase{
		Connection: connection,
	}

	return cockroachdbOneDatabase
}
