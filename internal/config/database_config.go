package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type DatabaseConfig struct {
	CockroachDatabase *CockroachDatabase
}

type CockroachDatabase struct {
	Pool *pgxpool.Pool
}

func NewDatabaseConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		CockroachDatabase: NewCockroachdbDatabase(envConfig),
	}
	return databaseConfig
}

func NewCockroachdbDatabase(envConfig *EnvConfig) *CockroachDatabase {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pool, poolErr := pgxpool.New(ctx, url)

	if poolErr != nil {
		panic(poolErr)
	}

	cockroachdbDatabase := &CockroachDatabase{
		Pool: pool,
	}

	return cockroachdbDatabase
}
