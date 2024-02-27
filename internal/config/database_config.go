package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	MariaDbOneDatabase *MariaDbOneDatabase
}

type MariaDbOneDatabase struct {
	Db *sql.DB
}

func NewDatabaseConfig(envConfig *EnvConfig) *DatabaseConfig {
	databaseConfig := &DatabaseConfig{
		MariaDbOneDatabase: NewMariaDbOneDatabase(envConfig),
	}
	return databaseConfig
}

func NewMariaDbOneDatabase(envConfig *EnvConfig) *MariaDbOneDatabase {
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		envConfig.MariadbOne.User,
		envConfig.MariadbOne.Password,
		envConfig.MariadbOne.Host,
		envConfig.MariadbOne.Port,
		envConfig.MariadbOne.Database,
	)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	return &MariaDbOneDatabase{
		Db: db,
	}
}
