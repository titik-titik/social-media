package http

import (
	"encoding/json"
	"net/http"
	model_request "social-media/internal/model/request/controller"
	model_response "social-media/internal/model/response"
	"social-media/internal/use_case"
)

type AuthController struct {
	AuthUseCase *use_case.AuthUseCase
}

func NewAuthController(authUseCase *use_case.AuthUseCase) *AuthController {
	AuthController := &AuthController{
		AuthUseCase: authUseCase,
	}
	return AuthController
}
func (authController *AuthController) Register(writer http.ResponseWriter, reader *http.Request) {
	request := &model_request.RegisterRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := authController.AuthUseCase.Register(request)
	model_response.NewResponse(writer, result.Message, result.Data, result.Code)
}

func (authController *AuthController) Login(writer http.ResponseWriter, reader *http.Request) {
	request := &model_request.LoginRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := authController.AuthUseCase.Login(request)
	model_response.NewResponse(writer, result.Message, result.Data, result.Code)
}
