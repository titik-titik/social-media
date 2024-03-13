package http

import (
	"encoding/json"
	"net/http"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/use_case"
	"strings"
)

type AuthController struct {
	AuthUseCase *use_case.AuthUseCase
}

func NewAuthController(authUseCase *use_case.AuthUseCase) *AuthController {
	return &AuthController{
		AuthUseCase: authUseCase,
	}
}

func (authController *AuthController) Register(writer http.ResponseWriter, reader *http.Request) {
	request := &model_request.RegisterRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := reader.Context()
	register, registerErr := authController.AuthUseCase.Register(ctx, request)
	if registerErr == nil {
		response.NewResponse(writer, register)
	}
}

func (authController *AuthController) Login(writer http.ResponseWriter, reader *http.Request) {
	request := &model_request.LoginRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx := reader.Context()
	login, loginErr := authController.AuthUseCase.Login(ctx, request)
	if loginErr == nil {
		response.NewResponse(writer, login)
	}
}

func (authController *AuthController) Logout(writer http.ResponseWriter, reader *http.Request) {
	token := reader.Header.Get("Authorization")
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	ctx := reader.Context()
	logout, logoutErr := authController.AuthUseCase.Logout(ctx, tokenString)
	if logoutErr == nil {
		response.NewResponse(writer, logout)
	}
}

func (authController *AuthController) GetNewAccessToken(writer http.ResponseWriter, reader *http.Request) {
	token := reader.Header.Get("Authorization")
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	ctx := reader.Context()
	getNewAccessToken, getNewAccessTokenErr := authController.AuthUseCase.GetNewAccessToken(ctx, tokenString)
	if getNewAccessTokenErr == nil {
		response.NewResponse(writer, getNewAccessToken)
	}
}
