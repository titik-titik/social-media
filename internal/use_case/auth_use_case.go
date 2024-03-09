package use_case

import (
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_controller "social-media/internal/model/request/controller"
	"social-media/internal/repository"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	DatabaseConfig    *config.DatabaseConfig
	UserRepository    *repository.UserRepository
	SessionRepository *repository.SessionRepository
}

func NewAuthUseCase(
	databaseConfig *config.DatabaseConfig,
	userRepository *repository.UserRepository,
	sessionRepository *repository.SessionRepository,
) *AuthUseCase {
	authUseCase := &AuthUseCase{
		DatabaseConfig:    databaseConfig,
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
	return authUseCase
}

func (authUseCase *AuthUseCase) Register(request *model_controller.RegisterRequest) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(request.Password.String), bcrypt.DefaultCost)
		if hashedPasswordErr != nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Register is failed, password hashing is failed.",
				Data:    nil,
			}
			return err
		}

		currentTime := null.NewTime(time.Now().UTC(), true)
		newUser := &entity.User{
			Id:         null.NewString(uuid.NewString(), true),
			Name:       request.Name,
			Username:   request.Username,
			Email:      request.Email,
			Password:   null.NewString(string(hashedPassword), true),
			AvatarUrl:  null.NewString("", false),
			Bio:        null.NewString("", false),
			IsVerified: null.NewBool(false, false),
			CreatedAt:  currentTime,
			UpdatedAt:  currentTime,
			DeletedAt:  null.NewTime(time.Time{}, false),
		}

		createdUser := authUseCase.UserRepository.CreateOne(begin, newUser)
		if createdUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Register is failed, user is not created.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusCreated,
			Message: "AuthUseCase Register is succeed.",
			Data:    createdUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Register is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (authUseCase *AuthUseCase) Login(request *model_controller.LoginRequest) (result *model.Result[*entity.Session]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		selectedUser := authUseCase.UserRepository.FindOneByEmail(begin, request.Email.String)
		if selectedUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "AuthUseCase Login is failed, user is not found by email.",
				Data:    nil,
			}
			return err
		}

		comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password.String), []byte(request.Password.String))
		if comparePasswordErr != nil {
			err = begin.Rollback()
			result = &model.Result[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "AuthUseCase Login is failed, password is not match.",
				Data:    nil,
			}
			return err
		}

		accessToken := null.NewString(uuid.NewString(), true)
		refreshToken := null.NewString(uuid.NewString(), true)
		currentTime := null.NewTime(time.Now().UTC(), true)
		accessTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Minute*10), true)
		refreshTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Hour*24*2), true)

		foundSession := authUseCase.SessionRepository.FindOneByUserId(begin, selectedUser.Id.String)
		if foundSession != nil {
			foundSession.AccessToken = accessToken
			foundSession.RefreshToken = refreshToken
			foundSession.AccessTokenExpiredAt = accessTokenExpiredAt
			foundSession.RefreshTokenExpiredAt = refreshTokenExpiredAt
			foundSession.UpdatedAt = currentTime
			updatedSession := authUseCase.SessionRepository.PatchOneById(begin, foundSession.Id.String, foundSession)

			if updatedSession == nil {
				err = begin.Rollback()
				result = &model.Result[*entity.Session]{
					Code:    http.StatusInternalServerError,
					Message: "AuthUseCase Login is failed, session is not patched.",
					Data:    nil,
				}
				return err
			}

			err = begin.Commit()
			result = &model.Result[*entity.Session]{
				Code:    http.StatusOK,
				Message: "AuthUseCase Login is succeed.",
				Data:    updatedSession,
			}
			return err
		}

		newSession := &entity.Session{
			Id:                    null.NewString(uuid.NewString(), true),
			UserId:                selectedUser.Id,
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiredAt:  accessTokenExpiredAt,
			RefreshTokenExpiredAt: refreshTokenExpiredAt,
			CreatedAt:             currentTime,
			UpdatedAt:             currentTime,
			DeletedAt:             null.NewTime(time.Time{}, false),
		}

		createdSession := authUseCase.SessionRepository.CreateOne(begin, newSession)
		if createdSession == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.Session]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Login is failed, session is not created.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.Session]{
			Code:    http.StatusCreated,
			Message: "AuthUseCase Login is succeed.",
			Data:    createdSession,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Login is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}
