package repository

import (
	"database/sql"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_repository "social-media/internal/model/request/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	DatabaseConfig *config.DatabaseConfig
}

func NewAuthRepository(databaseConfig *config.DatabaseConfig) *AuthRepository {
	authRepository := &AuthRepository{
		DatabaseConfig: databaseConfig,
	}
	return authRepository
}
func (authRepository *AuthRepository) Register(toRegisterUser *entity.User) *entity.User {
	begin, beginErr := authRepository.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		panic(beginErr)
	}

	rows, queryErr := begin.Query(
		"INSERT INTO \"user\" (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at;",
		toRegisterUser.Id,
		toRegisterUser.Name,
		toRegisterUser.Username,
		toRegisterUser.Email,
		toRegisterUser.Password,
		toRegisterUser.AvatarUrl,
		toRegisterUser.Bio,
		toRegisterUser.IsVerified,
		toRegisterUser.CreatedAt,
		toRegisterUser.UpdatedAt,
		toRegisterUser.DeletedAt,
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

func (authRepository *AuthRepository) Login(request *model_repository.LoginRepositoryRequest) *model.Result[*entity.Session] {
	begin, beginErr := authRepository.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "transaction error",
			Data:    nil,
		}
	}

	var user entity.User

	query := "SELECT id, username, email, password FROM \"user\" WHERE email = $1"

	err := begin.QueryRow(query, request.User.Email).Scan(
		&user.Id, &user.Username, &user.Email, &user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return &model.Result[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "invalid email or password not valid",
				Data:    nil,
			}
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(request.User.Password.String))
	if err == nil {
		return &model.Result[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "invalid email or password not valid",
			Data:    nil,
		}
	}
	rows := begin.QueryRow(
		"INSERT INTO \"session\" (id, user_id, access_token, access_token_expired_at, refresh_token, refresh_token_expired_at, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, user_id, access_token, access_token_expired_at, refresh_token, refresh_token_expired_at,  created_at, updated_at, deleted_at",
		request.Session.ID,
		user.Id,
		request.Session.AccessToken,
		request.Session.AccessTokenExpiredAt,
		request.Session.RefreshToken,
		request.Session.RefreshTokenExpiredAt,
		request.Session.CreatedAt,
		request.Session.UpdatedAt,
		request.Session.DeletedAt,
	)
	session := &entity.Session{}
	insertScanError := rows.Scan(
		session.ID,
		session.UserID,
		session.AccessToken,
		session.AccessTokenExpiredAt,
		session.RefreshToken,
		session.RefreshTokenExpiredAt,
		session.CreatedAt,
		session.UpdatedAt,
		session.DeletedAt,
	)
	if insertScanError != nil {
		return &model.Result[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "insert scan error",
			Data:    nil,
		}
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		return &model.Result[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "transaction error",
			Data:    nil,
		}
	}
	return &model.Result[*entity.Session]{
		Code:    http.StatusOK,
		Message: "login success",
		Data:    session,
	}

}
