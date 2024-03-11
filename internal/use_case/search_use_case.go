package use_case

import (
	"fmt"
	"net/http"
	"social-media/internal/config"
	model_controller "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/repository"
)

type SearchUseCase struct {
	DatabaseConfig *config.DatabaseConfig
	UserRepository *repository.UserRepository
}

func NewSearchUseCase(
	databaseConfig *config.DatabaseConfig,
	userRepository *repository.UserRepository,
) *SearchUseCase {
	searchUseCase := &SearchUseCase{
		DatabaseConfig: databaseConfig,
		UserRepository: userRepository,
	}
	return searchUseCase
}
func (sc *SearchUseCase) GetAllUsers(request *model_controller.GetAllUserRequest) *response.Response[[]*response.UserResponse] {
	tx, err := sc.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return &response.Response[[]*response.UserResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	users, err := sc.UserRepository.GetAllUsers(tx, request.Order, request.Limit, request.Offset)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return &response.Response[[]*response.UserResponse]{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("Failed to get all users: %+v, unable to rollback: %+v", err, rollbackErr),
			}
		}
		return &response.Response[[]*response.UserResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := tx.Commit(); err != nil {
		return &response.Response[[]*response.UserResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	userResponses := make([]*response.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = &response.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			AvatarUrl: user.AvatarUrl,
			// Tambahkan atribut lain sesuai kebutuhan
		}
	}

	return &response.Response[[]*response.UserResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    userResponses,
	}
}
