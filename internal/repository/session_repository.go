package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"social-media/internal/entity"
	"time"
)

type SessionRepository struct {
}

func NewSessionRepository() *SessionRepository {
	sessionRepository := &SessionRepository{}
	return sessionRepository
}

func DeserializeSessionRows(rows pgx.Rows) []*entity.Session {
	foundSessions, collectRowErr := pgx.CollectRows(rows, pgx.RowToStructByName[*entity.Session])
	if collectRowErr != nil {
		panic(collectRowErr)
	}
	return foundSessions
}

func (sessionRepository *SessionRepository) FindOneById(begin pgx.Tx, id string) *entity.Session {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at FROM \"session\" WHERE id=$1 LIMIT 1;",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		return nil
	}

	return foundSessions[0]
}

func (sessionRepository *SessionRepository) FindOneByUserId(begin pgx.Tx, userId string) *entity.Session {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"SELECT id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at FROM \"session\" WHERE user_id=$1 LIMIT 1;",
		userId,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		return nil
	}

	return foundSessions[0]
}

func (sessionRepository *SessionRepository) CreateOne(begin pgx.Tx, toCreateSession *entity.Session) *entity.Session {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, queryErr := begin.Query(
		ctx,
		"INSERT INTO \"session\" (id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);",
		toCreateSession.Id,
		toCreateSession.UserId,
		toCreateSession.AccessToken,
		toCreateSession.RefreshToken,
		toCreateSession.AccessTokenExpiredAt,
		toCreateSession.RefreshTokenExpiredAt,
		toCreateSession.CreatedAt,
		toCreateSession.UpdatedAt,
		toCreateSession.DeletedAt,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	return toCreateSession
}

func (sessionRepository *SessionRepository) PatchOneById(begin pgx.Tx, id string, toPatchSession *entity.Session) *entity.Session {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, queryErr := begin.Query(
		ctx,
		"UPDATE \"session\" SET id=$1, user_id=$2, access_token=$3, refresh_token=$4, access_token_expired_at=$5, refresh_token_expired_at=$6, created_at=$7, updated_at=$8, deleted_at=$9 WHERE id=$10 LIMIT 1;",
		toPatchSession.Id,
		toPatchSession.UserId,
		toPatchSession.AccessToken,
		toPatchSession.RefreshToken,
		toPatchSession.AccessTokenExpiredAt,
		toPatchSession.RefreshTokenExpiredAt,
		toPatchSession.CreatedAt,
		toPatchSession.UpdatedAt,
		toPatchSession.DeletedAt,
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	return toPatchSession
}

func (sessionRepository *SessionRepository) DeleteOneById(begin pgx.Tx, id string) *entity.Session {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, queryErr := begin.Query(
		ctx,
		"DELETE FROM \"session\" WHERE id=$1 LIMIT 1 RETURNING id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at;",
		id,
	)
	if queryErr != nil {
		panic(queryErr)
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		return nil
	}

	return foundSessions[0]
}
