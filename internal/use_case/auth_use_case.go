package use_case

import (
	"social-media/internal/repository"
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
