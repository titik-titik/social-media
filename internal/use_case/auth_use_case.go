package use_case

import (
	"context"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_controller "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
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

func (authUseCase *AuthUseCase) Register(ctx context.Context, request *model_controller.RegisterRequest) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(request.Password.String), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Register is failed, password hashing is failed.",
			Data:    nil,
		}
		err = nil
		return result, err
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

	createdUser, createdUserErr := authUseCase.UserRepository.CreateOne(transaction.Tx, newUser)
	if createdUserErr != nil {
		transaction.TxErr = createdUserErr
		result = nil
		err = createdUserErr
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusCreated,
		Message: "AuthUseCase Register is succeed.",
		Data:    createdUser,
	}
	err = nil
	return result, err
}

func (authUseCase *AuthUseCase) Login(ctx context.Context, request *model_controller.LoginRequest) (result *response.Response[*entity.Session], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundUser, foundSessionErr := authUseCase.UserRepository.FindOneByEmail(transaction.Tx, request.Email.String)
	if foundSessionErr != nil {
		transaction.TxErr = foundSessionErr
		result = nil
		err = foundSessionErr
		return result, err
	}

	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase Login is failed, user is not found by email.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(foundUser.Password.String), []byte(request.Password.String))
	if comparePasswordErr != nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase Login is failed, password is not match.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	accessToken := null.NewString(uuid.NewString(), true)
	refreshToken := null.NewString(uuid.NewString(), true)
	currentTime := null.NewTime(time.Now(), true)
	accessTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Minute*10), true)
	refreshTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Hour*24*2), true)

	foundSession, foundSessionErr := authUseCase.SessionRepository.FindOneByUserId(transaction.Tx, foundUser.Id.String)
	if foundSessionErr != nil {
		transaction.TxErr = foundSessionErr
		result = nil
		err = foundSessionErr
		return result, err
	}

	if foundSession != nil {
		foundSession.AccessToken = accessToken
		foundSession.RefreshToken = refreshToken
		foundSession.AccessTokenExpiredAt = accessTokenExpiredAt
		foundSession.RefreshTokenExpiredAt = refreshTokenExpiredAt
		foundSession.UpdatedAt = currentTime
		patchedSession, patchedSessionErr := authUseCase.SessionRepository.PatchOneById(transaction.Tx, foundSession.Id.String, foundSession)
		if patchedSessionErr != nil {
			transaction.TxErr = patchedSessionErr
			result = nil
			err = patchedSessionErr
			return result, err
		}

		result = &response.Response[*entity.Session]{
			Code:    http.StatusOK,
			Message: "AuthUseCase Login is succeed.",
			Data:    patchedSession,
		}
		err = nil
		return result, err
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

	createdSession, createdSessionErr := authUseCase.SessionRepository.CreateOne(transaction.Tx, newSession)
	if createdSessionErr != nil {
		transaction.TxErr = createdSessionErr
		result = nil
		err = createdSessionErr
		return result, err
	}

	result = &response.Response[*entity.Session]{
		Code:    http.StatusCreated,
		Message: "AuthUseCase Login is succeed.",
		Data:    createdSession,
	}
	err = nil
	return result, err
}

func (authUseCase *AuthUseCase) Logout(ctx context.Context, accessToken string) (result *response.Response[*entity.Session], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundSession, deletedSessionErr := authUseCase.SessionRepository.FindOneByAccToken(transaction.Tx, accessToken)
	if deletedSessionErr != nil {
		transaction.TxErr = deletedSessionErr
		result = nil
		err = deletedSessionErr
		return result, deletedSessionErr
	}

	if foundSession == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase Logout is failed, session is not found by access token.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	deletedSession, deletedSessionErr := authUseCase.SessionRepository.DeleteOneById(transaction.Tx, foundSession.Id.String)
	if deletedSessionErr != nil {
		transaction.TxErr = deletedSessionErr
		result = nil
		err = deletedSessionErr
		return result, deletedSessionErr
	}

	result = &response.Response[*entity.Session]{
		Code:    http.StatusOK,
		Message: "AuthUseCase Logout is succeed.",
		Data:    deletedSession,
	}
	err = nil
	return result, err
}

func (authUseCase *AuthUseCase) GetNewAccessToken(ctx context.Context, refreshToken string) (result *response.Response[*entity.Session], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundSession, patchedSessionErr := authUseCase.SessionRepository.FindOneByRefToken(transaction.Tx, refreshToken)
	if patchedSessionErr != nil {
		transaction.TxErr = patchedSessionErr
		result = nil
		err = patchedSessionErr
		return result, patchedSessionErr
	}

	if foundSession.RefreshTokenExpiredAt.Time.Before(time.Now()) {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase GetNewAccessToken is failed, refresh token is expired.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	foundSession.AccessToken = null.NewString(uuid.NewString(), true)
	foundSession.UpdatedAt = null.NewTime(time.Now(), true)
	patchedSession, patchedSessionErr := authUseCase.SessionRepository.PatchOneById(transaction.Tx, foundSession.Id.String, foundSession)
	if patchedSessionErr != nil {
		transaction.TxErr = patchedSessionErr
		result = nil
		err = patchedSessionErr
		return result, patchedSessionErr
	}
	result = &response.Response[*entity.Session]{
		Code:    http.StatusOK,
		Message: "AuthUseCase GetNewAccessToken is succeed.",
		Data:    patchedSession,
	}
	err = nil
	return result, err
}
