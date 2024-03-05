package config

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DatabaseConfig struct {
	CockroachdbDatabase *CockroachdbDatabase
}

type CockroachdbDatabase struct {
	Connection *sql.DB
}

func NewDatabaseConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		CockroachdbDatabase: NewCockroachdbDatabase(envConfig),
	}
	return databaseConfig
}

func NewCockroachdbDatabase(envConfig *EnvConfig) *CockroachdbDatabase {
	var url string
	if envConfig.Cockroachdb.Password == "" {
		url = fmt.Sprintf(
			"postgresql://%s@%s:%s/%s",
			envConfig.Cockroachdb.User,
			envConfig.Cockroachdb.Host,
			envConfig.Cockroachdb.Port,
			envConfig.Cockroachdb.Database,
		)
	} else {
		url = fmt.Sprintf(
			"postgresql://%s@%s:%s/%s",
			envConfig.Cockroachdb.User,
			envConfig.Cockroachdb.Host,
			envConfig.Cockroachdb.Port,
			envConfig.Cockroachdb.Database,
		)
	}

	connection, err := sql.Open(
		"pgx",
		url,
	)
	if err != nil {
		panic(err)
	}

	cockroachdbDatabase := &CockroachdbDatabase{
		Connection: connection,
	}

	return cockroachdbDatabase
}
