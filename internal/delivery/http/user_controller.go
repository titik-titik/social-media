package http

import (
	"encoding/json"
	"net/http"
	"social-media/internal/entity"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
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

	response.NewResponse(writer, result)
}

func (userController *UserController) FindOneByOneParam(writer http.ResponseWriter, reader *http.Request) {
	email := reader.URL.Query().Get("email")
	username := reader.URL.Query().Get("username")

	if email != "" {
		result := userController.UserUseCase.FindOneByEmail(email)
		response.NewResponse(writer, result)
	} else if username != "" {
		result := userController.UserUseCase.FindOneByUsername(username)
		response.NewResponse(writer, result)
	} else {
		result := &response.Response[*entity.User]{
			Message: "User parameter is invalid.",
			Data:    nil,
			Code:    http.StatusNotFound,
		}

		response.NewResponse(writer, result)
	}
}

func (userController *UserController) PatchOneById(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	request := &model_request.UserPatchOneByIdRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}

	result := userController.UserUseCase.PatchOneByIdFromRequest(id, request)

	response.NewResponse(writer, result)
}

func (userController *UserController) DeleteOneById(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	result := userController.UserUseCase.DeleteOneById(id)

	response.NewResponse(writer, result)
}
