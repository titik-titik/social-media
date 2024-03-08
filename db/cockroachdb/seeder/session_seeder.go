package seeder

import (
	"context"
	"social-media/internal/config"
	"social-media/test/mock"
	"time"
)

type SessionSeeder struct {
	DatabaseConfig *config.DatabaseConfig
	SessionMock    *mock.SessionMock
}

func NewSessionSeeder(
	databaseConfig *config.DatabaseConfig,
	userSeeder *UserSeeder,
) *SessionSeeder {
	sessionSeeder := &SessionSeeder{
		DatabaseConfig: databaseConfig,
		SessionMock:    mock.NewSessionMock(userSeeder.UserMock),
	}
	return sessionSeeder
}

func (sessionSeeder *SessionSeeder) Up() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, session := range sessionSeeder.SessionMock.Data {
		connection, acquireErr := sessionSeeder.DatabaseConfig.CockroachDatabase.Pool.Acquire(ctx)
		if acquireErr != nil {
			panic(acquireErr)
		}
		begin, beginErr := connection.Begin(ctx)
		if beginErr != nil {
			panic(beginErr)
		}

		_, queryErr := begin.Query(
			ctx,
			"INSERT INTO \"session\" (id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);",
			session.Id,
			session.UserId,
			session.AccessToken,
			session.RefreshToken,
			session.AccessTokenExpiredAt,
			session.RefreshTokenExpiredAt,
			session.CreatedAt,
			session.UpdatedAt,
			session.DeletedAt,
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

func (sessionSeeder *SessionSeeder) Down() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, session := range sessionSeeder.SessionMock.Data {
		connection, acquireErr := sessionSeeder.DatabaseConfig.CockroachDatabase.Pool.Acquire(ctx)
		if acquireErr != nil {
			panic(acquireErr)
		}
		begin, beginErr := connection.Begin(ctx)
		if beginErr != nil {
			panic(beginErr)
		}

		_, queryErr := begin.Query(
			ctx,
			"DELETE FROM \"session\" WHERE id = $1 LIMIT 1;",
			session.Id,
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
