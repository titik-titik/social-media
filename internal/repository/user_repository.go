package repository

import (
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

func (userRepository *UserRepository) FindOneById(id string) *entity.User {
	rows, queryErr := userRepository.Database.MariaDbOneDatabase.Db.Query(
		"SELECT id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at FROM user WHERE id = ? LIMIT 1", id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundUser := &entity.User{}

	for rows.Next() {
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
	}
	return foundUser
}

func (userRepository *UserRepository) FindOneByUsername(username string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByEmail(email string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByEmailAndPassword(email string, password string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByUsernameAndPassword(username string, password string) *entity.User {
	return nil
}

func (userRepository *UserRepository) CreateOne(toCreateUser *entity.User) *entity.User {
	return nil
}

func (userRepository *UserRepository) PatchOneById(id string, toPatchUser *entity.User) *entity.User {
	return nil
}

func (userRepository *UserRepository) DeleteOneById(id string) *entity.User {
	return nil
}
