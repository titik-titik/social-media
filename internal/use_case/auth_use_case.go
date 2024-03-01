package use_case

import "social-media/internal/repository"

type AuthUseCase struct {
	AuthRepository repository.AuthRepository
}

func NewAuthUseCase(AuthRepository repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		AuthRepository: AuthRepository,
	}
}
