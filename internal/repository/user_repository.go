package repository

import (
	"database/sql"
	"social-media/internal/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	userRepository := &UserRepository{}
	return userRepository
}

func DeserializeUserRows(rows *sql.Rows) []*entity.User {
	var foundUsers []*entity.User
	for rows.Next() {
		foundUser := &entity.User{}
		scanErr := rows.Scan(
			&foundUser.Id,
			&foundUser.Name,
			&foundUser.Username,
			&foundUser.Email,
			&foundUser.Password,
			&foundUser.AvatarUrl,
			&foundUser.Bio,
			&foundUser.IsVerified,
			&foundUser.CreatedAt,
			&foundUser.UpdatedAt,
			&foundUser.DeletedAt,
		)
		if scanErr != nil {
			panic(scanErr)
		}
		foundUsers = append(foundUsers, foundUser)
	}
	return foundUsers
}

func (userRepository *UserRepository) FindOneById(begin *sql.Tx, id string) *entity.User {
	rows, queryErr := begin.Query(
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

func (userRepository *UserRepository) FindOneByUsername(begin *sql.Tx, username string) *entity.User {
	rows, queryErr := begin.Query(
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

func (userRepository *UserRepository) FindOneByEmail(begin *sql.Tx, email string) *entity.User {
	rows, queryErr := begin.Query(
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

func (userRepository *UserRepository) FindOneByEmailAndPassword(begin *sql.Tx, email string, password string) *entity.User {
	rows, queryErr := begin.Query(
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

func (userRepository *UserRepository) FindOneByUsernameAndPassword(begin *sql.Tx, username string, password string) *entity.User {
	rows, queryErr := begin.Query(
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

func (userRepository *UserRepository) CreateOne(begin *sql.Tx, toCreateUser *entity.User) *entity.User {
	_, queryErr := begin.Query(
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

func (userRepository *UserRepository) PatchOneById(begin *sql.Tx, id string, toPatchUser *entity.User) *entity.User {
	_, queryErr := begin.Query(
		"UPDATE \"user\" SET id=$1, name=$2, username=$3, email=$4, password=$5, avatar_url=$6, bio=$7, is_verified=$8, created_at=$9, updated_at=$10, deleted_at=$11 WHERE id = $12;",
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

func (userRepository *UserRepository) DeleteOneById(begin *sql.Tx, id string) *entity.User {
	deletedUser, deletedUserQueryErr := begin.Query(
		"DELETE FROM \"user\" WHERE id=$1 LIMIT 1 RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at;",
		id,
	)
	if deletedUserQueryErr != nil {
		panic(deletedUserQueryErr)
	}

	foundUsers := DeserializeUserRows(deletedUser)
	if len(foundUsers) == 0 {
		return nil
	}

	return foundUsers[0]
}
