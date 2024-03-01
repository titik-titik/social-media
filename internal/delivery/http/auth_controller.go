package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-media/internal/model"
	"social-media/internal/model/response"
	"social-media/internal/use_case"
)

type AuthController struct {
	AuthUseCase *use_case.AuthUseCase
}

func NewAuthController(AuthUseCase *use_case.AuthUseCase) *AuthController {
	return &AuthController{AuthUseCase: AuthUseCase}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// Use NewErrorResponse for error response
		resp := response.NewResponse("Failed to read user data from the request", http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)

		responseJsonByte, marshalErr := json.Marshal(resp)
		if marshalErr != nil {
			panic(marshalErr)
		}
		_, writeErr := w.Write(responseJsonByte)
		if writeErr != nil {
			panic(writeErr)
		}
		return
	}

	err := c.AuthUseCase.Register(req.Username, req.Password, req.Email)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create user: %v", err)
		// Use NewErrorResponse for error response
		resp := response.NewResponse(errorMessage, http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Code)

		responseJsonByte, marshalErr := json.Marshal(resp)
		if marshalErr != nil {
			panic(marshalErr)
		}
		_, writeErr := w.Write(responseJsonByte)
		if writeErr != nil {
			panic(writeErr)
		}
		return
	}

	// Use NewSuccessResponse for success response
	resp := response.NewResponse("Success", http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	responseJsonByte, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		panic(marshalErr)
	}
	_, writeErr := w.Write(responseJsonByte)
	if writeErr != nil {
		panic(writeErr)
	}
}
