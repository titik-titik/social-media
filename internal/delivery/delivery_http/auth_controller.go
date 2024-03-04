package delivery_http

import (
	"social-media/internal/use_case"
)

type AuthController struct {
	AuthUseCase *use_case.AuthUseCase
}

func NewAuthController(AuthUseCase *use_case.AuthUseCase) *AuthController {
	AuthController := &AuthController{
		AuthUseCase: AuthUseCase,
	}
	return AuthController
}
