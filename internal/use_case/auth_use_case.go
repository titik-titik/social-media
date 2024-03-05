package use_case

import (
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
