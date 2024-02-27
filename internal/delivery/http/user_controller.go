package http

import (
	"encoding/json"
	"net/http"
	"social-media/internal/entity"
	"social-media/internal/model"
	"social-media/internal/model/request/user"
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

func (userController *UserController) FindOneByOneParam(writer http.ResponseWriter, reader *http.Request) {
	email := reader.URL.Query().Get("email")
	username := reader.URL.Query().Get("username")

	if email != "" {
		result := userController.UserUseCase.FindOneByEmail(email)
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
	} else if username != "" {
		result := userController.UserUseCase.FindOneByUsername(username)
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
	} else {
		response := &model.Response[*entity.User]{
			Message: "User parameter is invalid.",
			Data:    nil,
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		responseJsonByte, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			panic(marshalErr)
		}
		_, writeErr := writer.Write(responseJsonByte)
		if writeErr != nil {
			panic(writeErr)
		}
	}
}

func (userController *UserController) PatchOneById(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	request := &user.PatchOneById{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := userController.UserUseCase.PatchOneByIdFromRequest(id, request)
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

func (userController *UserController) DeleteOneById(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	result := userController.UserUseCase.DeleteOneById(id)
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
