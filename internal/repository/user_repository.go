package repository

import (
	"database/sql"
	"social-media/internal/config"
	"social-media/internal/entity"
)

type UserRepository struct {
	Database *config.DatabaseConfig
}

func NewUserRepository(database *config.DatabaseConfig) *UserRepository {
	userRepository := &UserRepository{
		Database: database,
	}
	return userRepository
}

func rowMapper(rows *sql.Rows) *entity.User {
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
	return foundUser
}

func (userRepository *UserRepository) FindOneById(id string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user WHERE id = ? LIMIT 1", id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var foundUser *entity.User = nil
	for rows.Next() {
		foundUser = rowMapper(rows)
	}
	return foundUser
}

func (userRepository *UserRepository) FindOneByUsername(username string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user WHERE username = ? LIMIT 1",
		username,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var foundUser *entity.User = nil
	for rows.Next() {
		foundUser = rowMapper(rows)
	}

	return foundUser
}

func (userRepository *UserRepository) FindOneByEmail(email string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user WHERE email = ? LIMIT 1",
		email,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var foundUser *entity.User = nil
	for rows.Next() {
		foundUser = rowMapper(rows)
	}

	return foundUser
}

func (userRepository *UserRepository) FindOneByEmailAndPassword(email string, password string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user WHERE email = ? AND password = ? LIMIT 1",
		email,
		password,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var foundUser *entity.User = nil
	for rows.Next() {
		foundUser = rowMapper(rows)
	}

	return foundUser
}

func (userRepository *UserRepository) FindOneByUsernameAndPassword(username string, password string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user WHERE username = ? AND password = ? LIMIT 1",
		username,
		password,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var foundUser *entity.User = nil
	for rows.Next() {
		foundUser = rowMapper(rows)
	}

	return foundUser
}

func (userRepository *UserRepository) CreateOne(toCreateUser *entity.User) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"INSERT INTO user (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at",
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

	var createdUser *entity.User = nil
	for rows.Next() {
		createdUser = rowMapper(rows)
	}

	return createdUser
}

func (userRepository *UserRepository) PatchOneById(id string, toPatchUser *entity.User) *entity.User {
	foundUser := userRepository.FindOneById(id)
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

	_, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"UPDATE user SET name = ?, username = ?, email = ?, password = ?, avatar_url = ?, bio = ?, is_verified = ?, created_at = ?, updated_at = ?, deleted_at = ? WHERE id = ?",
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

	return foundUser
}

func (userRepository *UserRepository) DeleteOneById(id string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"DELETE FROM user WHERE id = ? RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	var deletedUser *entity.User = nil
	for rows.Next() {
		deletedUser = rowMapper(rows)
	}

	return deletedUser
}
