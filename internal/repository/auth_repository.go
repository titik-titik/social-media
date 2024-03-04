package repository

import (
	"social-media/internal/config"
)

type AuthRepository struct {
	DatabaseConfig *config.DatabaseConfig
}

func NewAuthRepository(databaseConfig *config.DatabaseConfig) *AuthRepository {
	authRepository := &AuthRepository{
		DatabaseConfig: databaseConfig,
	}
	return authRepository
}
