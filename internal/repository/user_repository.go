package repository

import (
	"database/sql"
	"social-media/internal/config"
	"social-media/internal/entity"
)

type UserRepository struct {
	DatabaseConfig *config.DatabaseConfig
}

func NewUserRepository(databaseConfig *config.DatabaseConfig) *UserRepository {
	userRepository := &UserRepository{
		DatabaseConfig: databaseConfig,
	}
	return userRepository
}

func deserializeRows(rows *sql.Rows) []*entity.User {
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

		foundUser.CreatedAt.Time = foundUser.CreatedAt.Time.UTC()
		foundUser.UpdatedAt.Time = foundUser.UpdatedAt.Time.UTC()
		foundUser.DeletedAt.Time = foundUser.DeletedAt.Time.UTC()
		foundUsers = append(foundUsers, foundUser)
	}
	return foundUsers
}



func (userRepository *UserRepository) FindOneById(id string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE id=$1 LIMIT 1;",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := deserializeRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByUsername(username string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE username=$1 LIMIT 1;",
		username,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := deserializeRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByEmail(email string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE email=$1 LIMIT 1;",
		email,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := deserializeRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByEmailAndPassword(email string, password string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE email=$1 AND password=$2 LIMIT 1;",
		email,
		password,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := deserializeRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) FindOneByUsernameAndPassword(username string, password string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM \"user\" WHERE username=$1 AND password=$2 LIMIT 1;",
		username,
		password,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUsers := deserializeRows(rows)
	if len(foundUsers) == 0 {
		return nil
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUsers[0]
}

func (userRepository *UserRepository) CreateOne(toCreateUser *entity.User) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
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

func (userRepository *UserRepository) PatchOneById(id string, toPatchUser *entity.User) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
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

	foundUser := foundUsers[0]
	foundUser.Id = toPatchUser.Id
	foundUser.Name = toPatchUser.Name
	foundUser.Username = toPatchUser.Username
	foundUser.Email = toPatchUser.Email
	foundUser.Password = toPatchUser.Password
	foundUser.AvatarUrl = toPatchUser.AvatarUrl
	foundUser.Bio = toPatchUser.Bio
	foundUser.IsVerified = toPatchUser.IsVerified
	foundUser.CreatedAt = toPatchUser.CreatedAt
	foundUser.UpdatedAt = toPatchUser.UpdatedAt
	foundUser.DeletedAt = toPatchUser.DeletedAt

	_, queryErr := begin.Query(
		"UPDATE \"user\" SET id=$1, name=$2, username=$3, email=$4, password=$5, avatar_url=$6, bio=$7, is_verified=$8, created_at=$9, updated_at=$10, deleted_at=$11 WHERE id = $12;",
		foundUser.Id,
		foundUser.Name,
		foundUser.Username,
		foundUser.Email,
		foundUser.Password,
		foundUser.AvatarUrl,
		foundUser.Bio,
		foundUser.IsVerified,
		foundUser.CreatedAt,
		foundUser.UpdatedAt,
		foundUser.DeletedAt,
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		panic(commitErr)
	}

	return foundUser
}

func (userRepository *UserRepository) DeleteOneById(id string) *entity.User {
	begin, beginErr := userRepository.DatabaseConfig.PostgresOneDatabase.Connection.Begin()
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
