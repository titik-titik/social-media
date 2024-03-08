package seeder

import (
	"context"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"social-media/internal/config"
	"social-media/test/mock"
	"time"
)

type UserSeeder struct {
	DatabaseConfig *config.DatabaseConfig
	UserMock       *mock.UserMock
}

func NewUserSeeder(
	databaseConfig *config.DatabaseConfig,
) *UserSeeder {
	userSeeder := &UserSeeder{
		DatabaseConfig: databaseConfig,
		UserMock:       mock.NewUserMock(),
	}
	return userSeeder
}

func (userSeeder *UserSeeder) Up() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, user := range userSeeder.UserMock.Data {
		hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(user.Password.String), bcrypt.DefaultCost)
		if hashedPasswordErr != nil {
			panic(hashedPasswordErr)
		}
		password := null.NewString(string(hashedPassword), true)
		connection, acquireErr := userSeeder.DatabaseConfig.CockroachDatabase.Pool.Acquire(ctx)
		if acquireErr != nil {
			panic(acquireErr)
		}
		begin, beginErr := connection.Begin(ctx)
		if beginErr != nil {
			panic(beginErr)
		}
		_, err := begin.Query(
			ctx,
			"INSERT INTO \"user\" (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
			user.Id,
			user.Name,
			user.Username,
			user.Email,
			password,
			user.AvatarUrl,
			user.Bio,
			user.IsVerified,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt,
		)
		if err != nil {
			panic(err)
		}

		commitErr := begin.Commit(ctx)
		if commitErr != nil {
			panic(commitErr)
		}

		connection.Release()
	}
}

func (userSeeder *UserSeeder) Down() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, user := range userSeeder.UserMock.Data {
		connection, acquireErr := userSeeder.DatabaseConfig.CockroachDatabase.Pool.Acquire(ctx)
		if acquireErr != nil {
			panic(acquireErr)
		}
		begin, beginErr := connection.Begin(ctx)
		if beginErr != nil {
			panic(beginErr)
		}

		_, queryErr := begin.Query(
			ctx,
			"DELETE FROM \"user\" WHERE id = $1 LIMIT 1;",
			user.Id,
		)
		if queryErr != nil {
			panic(queryErr)
		}

		commitErr := begin.Commit(ctx)
		if commitErr != nil {
			panic(commitErr)
		}

		connection.Release()
	}
}
