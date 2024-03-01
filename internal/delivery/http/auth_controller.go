package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-media/internal/model/response"
	"social-media/internal/use_case"
	"time"
)

type AuthController struct {
	AuthUseCase *use_case.AuthUseCase
}

func NewAuthController(AuthUseCase *use_case.AuthUseCase) *AuthController {
	return &AuthController{AuthUseCase: AuthUseCase}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ErrorResponse(w, "Failed to read user data from the request", http.StatusBadRequest)
		return
	}

	err := c.AuthUseCase.Register(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create user: %v", err)
		response.ErrorResponse(w, errorMessage, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	// Create a user data object to be sent in the response
	userData := struct {
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}{
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	responses.SuccessResponse(w, "Success", userData, http.StatusCreated)
}
