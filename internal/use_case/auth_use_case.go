package use_case

import (
	"fmt"
	"net/http"
	"social-media/db/redis"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_request "social-media/internal/model/request"
	"social-media/internal/repository"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	AuthRepository *repository.AuthRepository
	RedisManager   *redis.RedisManager
}

func NewAuthUseCase(
	authRepository *repository.AuthRepository,
	redisManager *redis.RedisManager,
) *AuthUseCase {
	authUseCase := &AuthUseCase{
		AuthRepository: authRepository,
		RedisManager:   redisManager,
	}
	return authUseCase
}

func (authUseCase *AuthUseCase) Register(request *model_request.RegisterRequest) *model.Result[*entity.User] {
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
			Code:    500,
			Message: "AuthUseCase Register is failed, password hashing is failed.",
			Data:    nil,
		}
	}
	newUser.Password = null.NewString(string(hashedPassword), true)

	newUUID := uuid.New()

	newUser.Id = null.NewString(newUUID.String(), true)

	createdUser := authUseCase.AuthRepository.Register(newUser)
	if createdUser == nil {
		return &model.Result[*entity.User]{
			Code:    500,
			Message: "AuthUseCase Register is failed, user is not created.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "AuthUseCase Register is succeed.",
		Data:    createdUser,
	}
}
func (authUsecase *AuthUseCase) Login(request *model_request.LoginRequest) *model.TokenResult {
	if request.Email.String == "" || request.Password.String == "" {
		return model.NewTokenResult(http.StatusBadRequest, "Email and password must be provided", "")
	}

	user, err := authUsecase.AuthRepository.Login(request.Email.String)
	if err != nil {
		return model.NewTokenResult(http.StatusUnauthorized, fmt.Sprintf("User with email %s not found", request.Email.String), "")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(request.Password.String))
	if err != nil {
		return model.NewTokenResult(http.StatusUnauthorized, "Invalid login credentials", "")
	}

	accToken := fmt.Sprintf("%s:%s", user.Id.String, uuid.New().String())
	accExpiration := time.Now().Add(time.Minute * 15)
	accRedisKey := "access_token"

	err = authUsecase.RedisManager.InsertData(accRedisKey, []byte(accToken), accExpiration)
	if err != nil {
		return model.NewTokenResult(http.StatusInternalServerError, fmt.Sprintf("Failed to create access token: %v", err), "")
	}

	refToken := fmt.Sprintf("%s:%s", user.Id.String, uuid.New().String())
	refExpiration := time.Now().Add(time.Hour * 24 * 7)
	refRedisKey := "refresh_token"
	err = authUsecase.RedisManager.InsertData(refRedisKey, []byte(refToken), refExpiration)
	if err != nil {
		return model.NewTokenResult(http.StatusInternalServerError, fmt.Sprintf("Failed to create refresh token: %v", err), "")
	}
	return model.NewTokenResult(http.StatusOK, "Login successful", accToken)
}
