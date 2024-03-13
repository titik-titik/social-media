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

	ctx := reader.Context()
	foundUser, foundUserErr := userController.UserUseCase.FindOneById(ctx, id)
	if foundUserErr == nil {
		response.NewResponse(writer, foundUser)
	}
}

func (userController *UserController) FindOneByOneParam(writer http.ResponseWriter, reader *http.Request) {
	email := reader.URL.Query().Get("email")
	username := reader.URL.Query().Get("username")

	ctx := reader.Context()
	var foundUser *response.Response[*entity.User]
	var foundUserErr error
	if email != "" {
		foundUser, foundUserErr = userController.UserUseCase.FindOneByEmail(ctx, email)
	} else if username != "" {
		foundUser, foundUserErr = userController.UserUseCase.FindOneByUsername(ctx, username)
	} else {
		foundUser = &response.Response[*entity.User]{
			Message: "User parameter is invalid.",
			Data:    nil,
			Code:    http.StatusNotFound,
		}
	}

	if foundUserErr == nil {
		response.NewResponse(writer, foundUser)
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

	ctx := reader.Context()
	patchedUser, patchedUserErr := userController.UserUseCase.PatchOneByIdFromRequest(ctx, id, request)
	if patchedUserErr == nil {
		response.NewResponse(writer, patchedUser)
	}
}

func (userController *UserController) DeleteOneById(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	ctx := reader.Context()
	deletedUser, deletedUserErr := userController.UserUseCase.DeleteOneById(ctx, id)
	if deletedUserErr == nil {
		response.NewResponse(writer, deletedUser)
	}
}
