package repository

import (
	"database/sql"
	"fmt"
	"social-media/internal/config"
	"social-media/internal/entity"
	"time"
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
func (authRepository *AuthRepository) Register(toRegisterUser *entity.User) *entity.User {
	begin, beginErr := authRepository.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"INSERT INTO \"user\" (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at;",
		toRegisterUser.Id,
		toRegisterUser.Name,
		toRegisterUser.Username,
		toRegisterUser.Email,
		toRegisterUser.Password,
		toRegisterUser.AvatarUrl,
		toRegisterUser.Bio,
		toRegisterUser.IsVerified,
		toRegisterUser.CreatedAt,
		toRegisterUser.UpdatedAt,
		toRegisterUser.DeletedAt,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	createdUsers := deserializeRows(rows)
	if len(createdUsers) == 0 {
		return nil
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return createdUsers[0]
}
func (authRepository *AuthRepository) CheckUser(email string) (*entity.User, error) {
	var user entity.User

	query := "SELECT id, username, email, password FROM \"user\" WHERE email = $1"

	err := authRepository.DatabaseConfig.CockroachdbDatabase.Connection.QueryRow(query, email).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email %s", email)
		}
		return nil, err
	}

	return &user, nil
}
func (authRepository *AuthRepository) Login(id string, accToken string, accExpiration time.Time, refToken string, refExpiration time.Time) string {
	begin, beginErr := authRepository.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	_, queryErr := begin.Exec(
		"INSERT INTO \"session\" (id, access_token, access_token_expired_at, refresh_token, refresh_token_expired_at) VALUES ($1, $2, $3, $4, $5);",
		id,
		accToken,
		accExpiration,
		refToken,
		refExpiration,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return accToken
}
