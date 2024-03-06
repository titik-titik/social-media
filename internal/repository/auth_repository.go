package repository

import (
	"social-media/internal/config"
	"social-media/internal/entity"
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

func (userRepository *AuthRepository) CreateDummy(toCreateUser *entity.User) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"INSERT INTO \"user\" (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at;",
		toCreateUser.Id,
		toCreateUser.Name,
		toCreateUser.Username,
		toCreateUser.Email,
		toCreateUser.Password,
		toCreateUser.AvatarUrl,
		toCreateUser.Bio,
		toCreateUser.IsVerified,
		toCreateUser.CreatedAt,
		toCreateUser.UpdatedAt,
		toCreateUser.DeletedAt,
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

func (userRepository *AuthRepository) DeleteDummyById(id string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	foundRows, foundRowsErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE id=$1 LIMIT 1;",
		id,
	)
	if foundRowsErr != nil {
		panic(foundRowsErr)
	}

	foundUsers := deserializeRows(foundRows)
	if len(foundUsers) == 0 {
		return nil
	}

	_, queryErr := begin.Query(
		"DELETE FROM \"user\" WHERE id=$1 RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at;",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUsers[0]
}
