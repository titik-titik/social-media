package repository

import (
	"database/sql"
	"social-media/internal/entity"
)

type SessionRepository struct {
}

func NewSessionRepository() *SessionRepository {
	sessionRepository := &SessionRepository{}
	return sessionRepository
}

func DeserializeSessionRows(rows *sql.Rows) []*entity.Session {
	var foundSessions []*entity.Session
	for rows.Next() {
		foundSession := &entity.Session{}
		scanErr := rows.Scan(
			&foundSession.Id,
			&foundSession.UserId,
			&foundSession.AccessToken,
			&foundSession.RefreshToken,
			&foundSession.AccessTokenExpiredAt,
			&foundSession.RefreshTokenExpiredAt,
			&foundSession.CreatedAt,
			&foundSession.UpdatedAt,
			&foundSession.DeletedAt,
		)
		if scanErr != nil {
			panic(scanErr)
		}
		foundSessions = append(foundSessions, foundSession)
	}
	return foundSessions
}

func (sessionRepository *SessionRepository) FindOneById(begin *sql.Tx, id string) (result *entity.Session, err error) {
	rows, queryErr := begin.Query(
		"SELECT id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at FROM \"session\" WHERE id=$1 LIMIT 1;",
		id,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		result = nil
		err = nil
		return
	}

	result = foundSessions[0]
	err = nil
	return result, err
}

func (sessionRepository *SessionRepository) FindOneByUserId(begin *sql.Tx, userId string) (result *entity.Session, err error) {
	rows, queryErr := begin.Query(
		"SELECT id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at FROM \"session\" WHERE user_id=$1 LIMIT 1;",
		userId,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundSessions[0]
	err = nil
	return result, err
}

func (sessionRepository *SessionRepository) CreateOne(begin *sql.Tx, toCreateSession *entity.Session) (result *entity.Session, err error) {
	_, queryErr := begin.Query(
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
		result = nil
		err = queryErr
		return result, err
	}

	result = toCreateSession
	err = nil
	return result, err
}

func (sessionRepository *SessionRepository) PatchOneById(begin *sql.Tx, id string, toPatchSession *entity.Session) (result *entity.Session, err error) {
	_, queryErr := begin.Query(
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
		result = nil
		err = queryErr
		return result, err
	}

	result = toPatchSession
	err = nil
	return result, err
}

func (sessionRepository *SessionRepository) DeleteOneById(begin *sql.Tx, id string) (result *entity.Session, err error) {
	rows, queryErr := begin.Query(
		"DELETE FROM \"session\" WHERE id=$1 LIMIT 1 RETURNING id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at;",
		id,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundSessions[0]
	err = nil
	return result, err
}
func (sessionRepository *SessionRepository) FindOneByToken(begin *sql.Tx, accessToken string) (result *entity.Session, err error) {
	rows, queryErr := begin.Query(
		"SELECT id, user_id, access_token, refresh_token, access_token_expired_at, refresh_token_expired_at, created_at, updated_at, deleted_at FROM \"session\" WHERE access_token=$1 LIMIT 1;",
		accessToken,
	)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}

	foundSessions := DeserializeSessionRows(rows)
	if len(foundSessions) == 0 {
		result = nil
		err = nil
		return result, err
	}

	result = foundSessions[0]
	err = nil
	return result, err
}
