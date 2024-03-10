package use_case

import (
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	model_controller "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/repository"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb"

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

func (authUseCase *AuthUseCase) Register(request *model_controller.RegisterRequest) (result *response.Response[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result = nil
			return err
		}

		hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(request.Password.String), bcrypt.DefaultCost)
		if hashedPasswordErr != nil {
			err = begin.Rollback()
			result = &response.Response[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Register is failed, password hashing is failed.",
				Data:    nil,
			}
			return err
		}

		currentTime := null.NewTime(time.Now(), true)
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

		createdUser, err := authUseCase.UserRepository.CreateOne(begin, newUser)
		if err != nil {
			return err
		}

		err = begin.Commit()
		result = &response.Response[*entity.User]{
			Code:    http.StatusCreated,
			Message: "AuthUseCase Register is succeed.",
			Data:    createdUser,
		}
		return err
	})

	if beginErr != nil {
		result = &response.Response[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Register  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}

func (authUseCase *AuthUseCase) Login(request *model_controller.LoginRequest) (result *response.Response[*entity.Session]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result = nil
			return err
		}

		foundUser, err := authUseCase.UserRepository.FindOneByEmail(begin, request.Email.String)
		if err != nil {
			return err
		}

		if foundUser == nil {
			err = begin.Rollback()
			result = &response.Response[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "AuthUseCase Login is failed, user is not found by email.",
				Data:    nil,
			}
			return err
		}

		comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(foundUser.Password.String), []byte(request.Password.String))
		if comparePasswordErr != nil {
			err = begin.Rollback()
			result = &response.Response[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "AuthUseCase Login is failed, password is not match.",
				Data:    nil,
			}
			return err
		}

		accessToken := null.NewString(uuid.NewString(), true)
		refreshToken := null.NewString(uuid.NewString(), true)
		currentTime := null.NewTime(time.Now(), true)
		accessTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Minute*10), true)
		refreshTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Hour*24*2), true)

		foundSession, err := authUseCase.SessionRepository.FindOneByUserId(begin, foundUser.Id.String)
		if err != nil {
			return err
		}

		if foundSession != nil {
			foundSession.AccessToken = accessToken
			foundSession.RefreshToken = refreshToken
			foundSession.AccessTokenExpiredAt = accessTokenExpiredAt
			foundSession.RefreshTokenExpiredAt = refreshTokenExpiredAt
			foundSession.UpdatedAt = currentTime
			patchedSession, err := authUseCase.SessionRepository.PatchOneById(begin, foundSession.Id.String, foundSession)
			if err != nil {
				return err
			}

			err = begin.Commit()
			result = &response.Response[*entity.Session]{
				Code:    http.StatusOK,
				Message: "AuthUseCase Login is succeed.",
				Data:    patchedSession,
			}
			return err
		}

		newSession := &entity.Session{
			Id:                    null.NewString(uuid.NewString(), true),
			UserId:                foundUser.Id,
			AccessToken:           accessToken,
			RefreshToken:          refreshToken,
			AccessTokenExpiredAt:  accessTokenExpiredAt,
			RefreshTokenExpiredAt: refreshTokenExpiredAt,
			CreatedAt:             currentTime,
			UpdatedAt:             currentTime,
			DeletedAt:             null.NewTime(time.Time{}, false),
		}

		createdSession, err := authUseCase.SessionRepository.CreateOne(begin, newSession)
		if err != nil {
			return err
		}

		err = begin.Commit()
		result = &response.Response[*entity.Session]{
			Code:    http.StatusCreated,
			Message: "AuthUseCase Login is succeed.",
			Data:    createdSession,
		}
		return err
	})

	if beginErr != nil {
		result = &response.Response[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Login  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}
func (authUseCase *AuthUseCase) Logout(accessToken string) (result *response.Response[*entity.Session]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			return err
		}

		foundSession, err := authUseCase.SessionRepository.FindOneByAccToken(begin, accessToken)
		if err != nil {
			return err
		}

		if foundSession == nil {
			err = begin.Rollback()
			result = &response.Response[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "AuthUseCase Logout is failed, session is not found by access token.",
				Data:    nil,
			}
			return err
		}

		deletedSession, err := authUseCase.SessionRepository.DeleteOneById(begin, foundSession.Id.String)
		if err != nil {
			return err
		}

		err = begin.Commit()
		result = &response.Response[*entity.Session]{
			Code:    http.StatusOK,
			Message: "AuthUseCase Logout is succeed.",
			Data:    deletedSession,
		}
		return err
	})

	if beginErr != nil {
		result = &response.Response[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Logout  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}

func (authUseCase *AuthUseCase) GetNewAccessToken(refreshToken string) (result *response.Response[*entity.Session]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			return err
		}

		foundSession, err := authUseCase.SessionRepository.FindOneByRefToken(begin, refreshToken)
		if err != nil {
			return err
		}

		if foundSession.RefreshTokenExpiredAt.Time.Before(time.Now()) {
			err = begin.Rollback()
			result = &response.Response[*entity.Session]{
				Code:    http.StatusNotFound,
				Message: "AuthUseCase GetNewAccessToken is failed, refresh token is expired.",
				Data:    nil,
			}
			return err

		}

		foundSession.AccessToken = null.NewString(uuid.NewString(), true)
		foundSession.UpdatedAt = null.NewTime(time.Now(), true)
		patchedSession, err := authUseCase.SessionRepository.PatchOneById(begin, foundSession.Id.String, foundSession)
		if err != nil {
			return err
		}

		err = begin.Commit()
		result = &response.Response[*entity.Session]{
			Code:    http.StatusOK,
			Message: "AuthUseCase GetNewAccessToken is succeed.",
			Data:    patchedSession,
		}
		return nil
	})

	if beginErr != nil {
		result = &response.Response[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase GetNewAccessToken is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}
