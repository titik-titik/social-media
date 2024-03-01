package repository

import (
	"social-media/internal/config"
	"time"
)

type AuthRepository struct {
	DatabaseConfig *config.DatabaseConfig
	//PostgresOneDB  *config.PostgresOneDatabase
}

func NewAuthRepository(databaseConfig *config.DatabaseConfig) *AuthRepository {
	authRepository := &AuthRepository{
		DatabaseConfig: databaseConfig,
		//PostgresOneDB:  databaseConfig.PostgresOneDatabase,
	}
	return authRepository
}

func (ar *AuthRepository) Register(id, username, password, email, avatarURL, bio string, isVerified bool, createdAt, updatedAt time.Time) error {
	//createSQL := `
	//    INSERT INTO users (id, username, password, email, avatar_url, bio, is_verified, created_at, updated_at)
	//    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	//`
	//
	//_, err := ar.PostgresOneDB.Connection.Exec(createSQL, id, username, password, email, avatarURL, bio, isVerified, createdAt, updatedAt)
	//return err
	return nil
}
