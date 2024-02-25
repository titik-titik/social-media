package http

import (
	"social-media/internal/use_case"
)

type UserController struct {
	UserUseCase *use_case.UserUseCase
}

func NewUserController(userUseCase *use_case.UserUseCase) *UserController {
	userController := &UserController{
		UserUseCase: userUseCase,
	}
	return userController
}
