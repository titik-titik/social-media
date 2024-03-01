package http

import (
	"social-media/internal/use_case"
)

type AuthController struct {
	UseCase *use_case.AuthUseCase
}

func NewAuthController(useCase *use_case.AuthUseCase) *AuthController {
	return &AuthController{UseCase: useCase}
}
