package use_case

import (
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

func (authUseCase *AuthUseCase) Register(request *model_controller.RegisterRequest) *model.Result[*entity.User] {
	begin, beginErr := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Register is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	newUser := &entity.User{
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		CreatedAt: null.NewTime(time.Now().UTC(), true),
		UpdatedAt: null.NewTime(time.Now().UTC(), true),
	}

	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(request.Password.String), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Register is failed, password hashing is failed.",
			Data:    nil,
		}
	}
	newUser.Password = null.NewString(string(hashedPassword), true)

	newUUId := uuid.New()

	newUser.Id = null.NewString(newUUId.String(), true)

	createdUser := authUseCase.UserRepository.CreateOne(begin, newUser)
	if createdUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Register is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Register is failed, user is not created.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusCreated,
		Message: "AuthUseCase Register is succeed.",
		Data:    createdUser,
	}
}

func (authUseCase *AuthUseCase) Login(request *model_controller.LoginRequest) *model.Result[*entity.Session] {
	begin, beginErr := authUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Login is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	selectedUser := authUseCase.UserRepository.FindOneByEmail(begin, request.Email.String)
	if selectedUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.Session]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Login is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase Login is failed, user is not found by email.",
			Data:    nil,
		}
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password.String), []byte(request.Password.String))
	if comparePasswordErr != nil {
		return &model.Result[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase Login is failed, password is not match.",
			Data:    nil,
		}
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
			rollbackEr := begin.Rollback()
			if rollbackEr != nil {
				return &model.Result[*entity.Session]{
					Code:    http.StatusInternalServerError,
					Message: "AuthUseCase Login is failed, transaction rollback is failed.",
					Data:    nil,
				}
			}
			return &model.Result[*entity.Session]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Login is failed, session is not patched.",
				Data:    nil,
			}
		}

		return &model.Result[*entity.Session]{
			Code:    http.StatusOK,
			Message: "AuthUseCase Login is succeed.",
			Data:    updatedSession,
		}
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
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.Session]{
				Code:    http.StatusInternalServerError,
				Message: "AuthUseCase Login is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Login is failed, session is not created.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.Session]{
		Code:    http.StatusCreated,
		Message: "AuthUseCase Login is succeed.",
		Data:    createdSession,
	}
}
