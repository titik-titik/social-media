package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	model_request "social-media/internal/model/request/controller"
	model_response "social-media/internal/model/response"
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

	result := authController.AuthUseCase.Register(request)
	model_response.NewResponse(writer, result.Message, result.Data, result.Code)
}

func (authController *AuthController) Login(writer http.ResponseWriter, reader *http.Request) {
	request := &model_request.LoginRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := authController.AuthUseCase.Login(request)
	model_response.NewResponse(writer, result.Message, result.Data, result.Code)
}

func (authController *AuthController) Logout(writer http.ResponseWriter, reader *http.Request) {
	token := reader.Header.Get("Authorization")
	tokenString := strings.Replace(token, "Bearer ", "", 1)
	fmt.Println(tokenString)

	result := authController.AuthUseCase.Logout(tokenString)
	model_response.NewResponse(writer, result.Message, result.Data, result.Code)
}

func (authController *AuthController) GetNewAccessToken(writer http.ResponseWriter, reader *http.Request) {
	token := reader.Header.Get("Authorization")
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	result := authController.AuthUseCase.GetNewAccessToken(tokenString)
	model_response.NewResponse(writer, result.Message, result.Data, result.Code)
}
