package use_case

import (
	"github.com/guregu/null"
	"social-media/internal/entity"
	"social-media/internal/model"
	"social-media/internal/model/request/user"
	"social-media/internal/repository"
	"time"
)

type UserUseCase struct {
	UserRepository *repository.UserRepository
}



func NewUserUseCase(
	userRepository *repository.UserRepository,
) *UserUseCase {
	userUseCase := &UserUseCase{
		UserRepository: userRepository,
	}
	return userUseCase
}

func (userUseCase *UserUseCase) FindOneById(id string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneById(id)

	if foundUser == nil {
		return &model.Result[*entity.User]{
			Code:    404,
			Message: "UserUserCase FindOneById is failed, user is not found by id.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneById is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByUsername(username string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByUsername(username)

	if foundUser == nil {
		return &model.Result[*entity.User]{
			Code:    404,
			Message: "UserUserCase FindOneByUsername is failed, user is not found by username.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByUsername is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByEmail(email string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByEmail(email)

	if foundUser == nil {
		return &model.Result[*entity.User]{
			Code:    404,
			Message: "UserUserCase FindOneByEmail is failed, user is not found by email.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByEmail is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByEmailAndPassword(email, password string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByEmailAndPassword(email, password)

	if foundUser == nil {
		return &model.Result[*entity.User]{
			Code:    404,
			Message: "UserUserCase FindOneByEmailAndPassword is failed, user is not found by email and password.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByEmailAndPassword is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByUsernameAndPassword(username, password string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByUsernameAndPassword(username, password)

	if foundUser == nil {
		return &model.Result[*entity.User]{
			Code:    404,
			Message: "UserUserCase FindOneByUsernameAndPassword is failed, user is not found by username and password.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByUsernameAndPassword is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) CreateOne(toCreateUser *entity.User) *model.Result[*entity.User] {
	createdUser := userUseCase.UserRepository.CreateOne(toCreateUser)

	if createdUser == nil {
		return &model.Result[*entity.User]{
			Code:    500,
			Message: "UserUserCase CreateOne is failed, user is not created.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase CreateOne is succeed.",
		Data:    createdUser,
	}
}

func (userUseCase *UserUseCase) PatchOneById(id string, toPatchUser *entity.User) *model.Result[*entity.User] {
	patchedUser := userUseCase.UserRepository.PatchOneById(id, toPatchUser)

	if patchedUser == nil {
		return &model.Result[*entity.User]{
			Code:    500,
			Message: "UserUserCase PatchOneById is failed, user is not patched.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase PatchOneById is succeed.",
		Data:    patchedUser,
	}
}

func (userUseCase *UserUseCase) PatchOneByIdFromRequest(id string, request *user.PatchOneById) *model.Result[*entity.User] {
	foundUserResult := userUseCase.FindOneById(id)
	if foundUserResult.Code != 200 || foundUserResult.Data == nil {
		return &model.Result[*entity.User]{
			Code:    404,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, user is not found by id.",
			Data:    nil,
		}
	}

	foundUserResult.Data.Name = request.Name
	foundUserResult.Data.Username = request.Username
	foundUserResult.Data.Email = request.Email
	foundUserResult.Data.Password = request.Password
	foundUserResult.Data.AvatarUrl = request.AvatarUrl
	foundUserResult.Data.Bio = request.Bio
	foundUserResult.Data.UpdatedAt = null.NewTime(time.Now().UTC(), true)

	patchedUserResult := userUseCase.PatchOneById(id, foundUserResult.Data)
	if patchedUserResult.Code != 200 || patchedUserResult.Data == nil {
		return &model.Result[*entity.User]{
			Code:    500,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, user is not patched.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase PatchOneByIdFromRequest is succeed.",
		Data:    patchedUserResult.Data,
	}
}

func (userUseCase *UserUseCase) DeleteOneById(id string) *model.Result[*entity.User] {
	deletedUser := userUseCase.UserRepository.DeleteOneById(id)

	if deletedUser == nil {
		return &model.Result[*entity.User]{
			Code:    500,
			Message: "UserUserCase DeleteOneById is failed, user is not deleted.",
			Data:    nil,
		}

	}

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase DeleteOneById is succeed.",
		Data:    deletedUser,
	}
}
