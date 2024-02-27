package http

import (
	"encoding/json"
	"net/http"
	"social-media/internal/model"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
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

func (userController *UserController) FindOneById(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	result := userController.UserUseCase.FindOneById(id)
	response := model.NewResponse(result.Message, result.Data)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(result.Code)
	responseJsonByte, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		panic(marshalErr)
	}
	_, writeErr := writer.Write(responseJsonByte)
	if writeErr != nil {
		panic(writeErr)
	}
}
