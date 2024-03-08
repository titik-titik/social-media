package use_case

import (
	"net/http"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_controller "social-media/internal/model/request/controller"
	model_repository "social-media/internal/model/request/repository"
	"social-media/internal/repository"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	AuthRepository *repository.AuthRepository
}

func NewAuthUseCase(
	authRepository *repository.AuthRepository,
) *AuthUseCase {
	authUseCase := &AuthUseCase{
		AuthRepository: authRepository,
	}
	return authUseCase
}

func (authUseCase *AuthUseCase) Register(request *model_controller.RegisterRequest) *model.Result[*entity.User] {
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
		Code:    http.StatusCreated,
		Message: "AuthUseCase Register is succeed.",
		Data:    createdUser,
	}
}

func (ac *AuthUseCase) Login(request *model_controller.LoginRequest) *model.Result[*entity.Session] {
	id := uuid.New().String()

	accToken := uuid.New().String()
	if accToken == "" {
		return &model.Result[*entity.Session]{
			Code:    500,
			Message: "AuthUseCase Register is failed, password hashing is failed.",
			Data:    nil,
		}
	}
	accExpiration := time.Now().Add(time.Minute * 15)

	refToken := uuid.New().String()
	if refToken == "" {
		return &model.Result[*entity.Session]{
			Code:    500,
			Message: "AuthUseCase Register is failed, password hashing is failed.",
			Data:    nil,
		}
	}
	currentTime := time.Now()
	refExpiration := time.Now().Add(time.Hour * 24 * 7)
	repositoryRequest := &model_repository.LoginRepositoryRequest{
		LoginControllerRequest: request,
		Session: &entity.Session{
			ID:                    null.NewString(id, true),
			AccessToken:           null.NewString(accToken, true),
			RefreshToken:          null.NewString(refToken, true),
			AccessTokenExpiredAt:  null.NewTime(accExpiration, true),
			RefreshTokenExpiredAt: null.NewTime(refExpiration, true),
			CreatedAt:             null.NewTime(currentTime, true),
			UpdatedAt:             null.NewTime(currentTime, true),
			DeletedAt:             null.NewTime(time.Time{}, false),
		},
		User: &entity.User{
			Email:    request.Email,
			Password: request.Password,
		},
	}

	result := ac.AuthRepository.Login(repositoryRequest)
	if result == nil {
		return &model.Result[*entity.Session]{
			Code:    500,
			Message: "AuthUseCase Register is failed, user is not created.",
			Data:    nil,
		}
	}
	return &model.Result[*entity.Session]{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data:    result.Data,
	}
}
