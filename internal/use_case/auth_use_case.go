package use_case

import (
	"fmt"
	"net/http"
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
}

func NewAuthUseCase(
	authRepository *repository.AuthRepository,
) *AuthUseCase {
	authUseCase := &AuthUseCase{
		AuthRepository: authRepository,
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

func (ac *AuthUseCase) Login(request *model_request.LoginRequest) *model.TokenResult {
	if request.Email.String == "" || request.Password.String == "" {
		return model.NewTokenResult(http.StatusBadRequest, "Email and password must be provided", "")
	}
	user, err := ac.AuthRepository.CheckUser(request.Email.String)
	if err != nil {
		return model.NewTokenResult(http.StatusUnauthorized, fmt.Sprintf("User with email %s not found", request.Email.String), "")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(request.Password.String))
	if err != nil {
		return model.NewTokenResult(http.StatusUnauthorized, "Invalid login credentials", "")
	}

	accToken := fmt.Sprintf("%s:%s", user.Id.String, uuid.New().String())
	if accToken == "" {
		errMsg := "Failed to create access token"
		return model.NewTokenResult(http.StatusInternalServerError, errMsg, "")
	}
	accExpiration := time.Now().Add(time.Minute * 15)

	refToken := fmt.Sprintf("%s:%s", user.Id.String, uuid.New().String())
	if refToken == "" {
		errMsg := "Failed to create access token"
		return model.NewTokenResult(http.StatusInternalServerError, errMsg, "")
	}
	refExpiration := time.Now().Add(time.Hour * 24 * 7)
	id := uuid.New().String()
	loginUser := ac.AuthRepository.Login(id, accToken, accExpiration, refToken, refExpiration)

	return model.NewTokenResult(http.StatusOK, "Login successful", loginUser)
}
