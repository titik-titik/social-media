package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"social-media/internal/entity"
	"time"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	userRepository := &UserRepository{}
	return userRepository
}

func DeserializeUserRows(rows pgx.Rows) []*entity.User {
	foundUsers, collectRowErr := pgx.CollectRows(rows, pgx.RowToStructByName[*entity.User])
	if collectRowErr != nil {
		panic(collectRowErr)
	}
	return foundUsers
}

func (userRepository *UserRepository) FindOneById(begin pgx.Tx, id string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE id=$1 LIMIT 1;",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByUsername(begin pgx.Tx, username string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE username=$1 LIMIT 1;",
		username,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByEmail(begin pgx.Tx, email string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE email=$1 LIMIT 1;",
		email,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByEmailAndPassword(begin pgx.Tx, email string, password string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE email=$1 AND password=$2 LIMIT 1;",
		email,
		password,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByUsernameAndPassword(begin pgx.Tx, username string, password string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE username=$1 AND password=$2 LIMIT 1;",
		username,
		password,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) CreateOne(begin pgx.Tx, toCreateUser *entity.User) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, queryErr := begin.Query(
		ctx,
		"INSERT INTO \"user\" (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
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

	return toCreateUser
}

func (userRepository *UserRepository) PatchOneById(begin pgx.Tx, id string, toPatchUser *entity.User) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, queryErr := begin.Query(
		ctx,
		"UPDATE \"user\" SET id=$1, name=$2, username=$3, email=$4, password=$5, avatar_url=$6, bio=$7, is_verified=$8, created_at=$9, updated_at=$10, deleted_at=$11 WHERE id = $12 LIMIT 1;",
		toPatchUser.Id,
		toPatchUser.Name,
		toPatchUser.Username,
		toPatchUser.Email,
		toPatchUser.Password,
		toPatchUser.AvatarUrl,
		toPatchUser.Bio,
		toPatchUser.IsVerified,
		toPatchUser.CreatedAt,
		toPatchUser.UpdatedAt,
		toPatchUser.DeletedAt,
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	return toPatchUser
}

func (userRepository *UserRepository) DeleteOneById(begin pgx.Tx, id string) *entity.User {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"DELETE FROM \"user\" WHERE id=$1 LIMIT 1 RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at;",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := DeserializeUserRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}
