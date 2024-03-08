package seeder

import (
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"social-media/internal/config"
	"social-media/test/mock"
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
	for _, session := range sessionSeeder.SessionMock.Data {
		begin, beginErr := sessionSeeder.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if beginErr != nil {
			panic(beginErr)
		}

		queryErr := crdb.Execute(func() (err error) {
			_, err = begin.Query(
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
			return err
		})
		if queryErr != nil {
			panic(queryErr)
		}

		commitErr := crdb.Execute(func() (err error) {
			err = begin.Commit()
			return err
		})
		if commitErr != nil {
			panic(commitErr)
		}
	}
}

func (sessionSeeder *SessionSeeder) Down() {
	for _, session := range sessionSeeder.SessionMock.Data {
		begin, beginErr := sessionSeeder.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if beginErr != nil {
			panic(beginErr)
		}

		queryErr := crdb.Execute(func() (err error) {
			_, err = begin.Query(
				"DELETE FROM \"session\" WHERE id=$1;",
				session.Id,
			)
			return err
		})
		if queryErr != nil {
			panic(queryErr)
		}

		commitErr := crdb.Execute(func() (err error) {
			err = begin.Commit()
			return err
		})
		if commitErr != nil {
			panic(commitErr)
		}
	}
}
