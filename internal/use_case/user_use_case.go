package use_case

import (
	"social-media/internal/entity"
	"social-media/internal/model"
	"social-media/internal/repository"
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

	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneById succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByUsername(username string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByUsername(username)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByUsername succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByEmail(email string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByEmail(email)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByEmail succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByEmailAndPassword(email, password string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByEmailAndPassword(email, password)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByEmailAndPassword succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByUsernameAndPassword(username, password string) *model.Result[*entity.User] {
	foundUser := userUseCase.UserRepository.FindOneByUsernameAndPassword(username, password)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase FindOneByUsernameAndPassword succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) CreateOne(toCreateUser *entity.User) *model.Result[*entity.User] {
	createdUser := userUseCase.UserRepository.CreateOne(toCreateUser)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase CreateOne succeed.",
		Data:    createdUser,
	}
}

func (userUseCase *UserUseCase) PatchOneById(id string, toPatchUser *entity.User) *model.Result[*entity.User] {
	patchedUser := userUseCase.UserRepository.PatchOneById(id, toPatchUser)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase PatchOneById succeed.",
		Data:    patchedUser,
	}
}

func (userUseCase *UserUseCase) DeleteOneById(id string) *model.Result[*entity.User] {
	deletedUser := userUseCase.UserRepository.DeleteOneById(id)
	return &model.Result[*entity.User]{
		Code:    200,
		Message: "UserUserCase DeleteOneById succeed.",
		Data:    deletedUser,
	}
}
