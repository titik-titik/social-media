package http

//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"social-media/internal/model"
//	"social-media/internal/model/response"
//	"social-media/internal/use_case"
//)
//
//type AuthController struct {
//	AuthUseCase *use_case.AuthUseCase
//}
//
//func NewAuthController(AuthUseCase *use_case.AuthUseCase) *AuthController {
//	return &AuthController{AuthUseCase: AuthUseCase}
//}
//
//func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
//	var req model.RegisterRequest
//
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		response.NewResponse(w, "Failed to read user data from the request", []string{}, http.StatusBadRequest)
//		return
//	}
//
//	err := c.AuthUseCase.Register(req.Username, req.Password, req.Email)
//	if err != nil {
//		errorMessage := fmt.Sprintf("Failed to create user: %v", err)
//		response.NewResponse(w, errorMessage, []string{}, http.StatusInternalServerError)
//		return
//	}
//
//	response.NewResponse(w, "Success", []string{}, http.StatusCreated)
//
//}
