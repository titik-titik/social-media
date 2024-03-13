package use_case

import (
	"context"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/repository"
	"time"
)

type UserUseCase struct {
	DatabaseConfig    *config.DatabaseConfig
	UserRepository    *repository.UserRepository
	SessionRepository *repository.SessionRepository
	PostRepository    *repository.PostRepository
}

func NewUserUseCase(
	databaseConfig *config.DatabaseConfig,
	userRepository *repository.UserRepository,
	sessionRepository *repository.SessionRepository,
	postRepository *repository.PostRepository,
) *UserUseCase {
	userUseCase := &UserUseCase{
		DatabaseConfig:    databaseConfig,
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
		PostRepository:    postRepository,
	}
	return userUseCase
}

func (userUseCase *UserUseCase) FindOneById(ctx context.Context, id string) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundUser, foundUserErr := userUseCase.UserRepository.FindOneById(transaction.Tx, id)
	if foundUserErr != nil {
		transaction.TxErr = foundUserErr
		result = nil
		err = foundUserErr
		return result, err
	}
	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneById is failed, user is not found by id.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneById is succeed.",
		Data:    foundUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) FindOneByUsername(ctx context.Context, username string) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)
	foundUser, foundUserErr := userUseCase.UserRepository.FindOneByUsername(transaction.Tx, username)
	if foundUserErr != nil {
		transaction.TxErr = foundUserErr
		result = nil
		err = foundUserErr
		return result, err
	}
	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByUsername is failed, user is not found by username.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByUsername is succeed.",
		Data:    foundUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) FindOneByEmail(ctx context.Context, email string) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundUser, foundUserErr := userUseCase.UserRepository.FindOneByEmail(transaction.Tx, email)
	if foundUserErr != nil {
		transaction.TxErr = foundUserErr
		result = nil
		err = foundUserErr
		return result, err
	}
	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByEmail is failed, user is not found by email.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByEmail is succeed.",
		Data:    foundUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) FindOneByEmailAndPassword(ctx context.Context, email, password string) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundUser, foundUserErr := userUseCase.UserRepository.FindOneByEmailAndPassword(transaction.Tx, email, password)
	if foundUserErr != nil {
		transaction.TxErr = foundUserErr
		result = nil
		return result, foundUserErr
	}
	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByEmailAndPassword is failed, user is not found by email and password.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByEmailAndPassword is succeed.",
		Data:    foundUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) FindOneByUsernameAndPassword(ctx context.Context, username, password string) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundUser, foundUserErr := userUseCase.UserRepository.FindOneByUsernameAndPassword(transaction.Tx, username, password)
	if foundUserErr != nil {
		transaction.TxErr = foundUserErr
		result = nil
		err = foundUserErr
		return result, err
	}
	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByUsernameAndPassword is failed, user is not found by username and password.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByUsernameAndPassword is succeed.",
		Data:    foundUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) CreateOne(ctx context.Context, toCreateUser *entity.User) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	createdUser, createdUserErr := userUseCase.UserRepository.CreateOne(transaction.Tx, toCreateUser)
	if createdUserErr != nil {
		transaction.TxErr = createdUserErr
		result = nil
		err = createdUserErr
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase CreateOne is succeed.",
		Data:    createdUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) PatchOneById(ctx context.Context, id string, toPatchUser *entity.User) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	patchedUser, patchedUserErr := userUseCase.UserRepository.PatchOneById(transaction.Tx, id, toPatchUser)
	if patchedUserErr != nil {
		transaction.TxErr = patchedUserErr
		result = nil
		err = patchedUserErr
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase UserPatchOneById is succeed.",
		Data:    patchedUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) PatchOneByIdFromRequest(ctx context.Context, id string, request *model_request.UserPatchOneByIdRequest) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	foundUser, foundUserErr := userUseCase.UserRepository.FindOneById(transaction.Tx, id)
	if foundUserErr != nil {
		transaction.TxErr = foundUserErr
		result = nil
		err = foundUserErr
		return result, err
	}
	if foundUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase PatchOneById is failed, user is not found by id.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	if request.Name.Valid {
		foundUser.Name = request.Name
	}
	if request.Email.Valid {
		foundUser.Email = request.Email
	}
	if request.Username.Valid {
		foundUser.Username = request.Username
	}
	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(request.Password.String), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		result = &response.Response[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, password hashing is failed.",
			Data:    nil,
		}
		err = nil
		return result, err
	}
	if request.Password.Valid {
		foundUser.Password = null.NewString(string(hashedPassword), true)
	}
	if request.AvatarUrl.Valid {
		foundUser.AvatarUrl = request.AvatarUrl
	}
	if request.Bio.Valid {
		foundUser.Bio = request.Bio
	}

	foundUser.UpdatedAt = null.NewTime(time.Now(), true)

	patchedUser, patchedUserErr := userUseCase.UserRepository.PatchOneById(transaction.Tx, id, foundUser)
	if patchedUserErr != nil {
		transaction.TxErr = patchedUserErr
		result = nil
		err = patchedUserErr
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase PatchOneByIdFromRequest is succeed.",
		Data:    patchedUser,
	}
	err = nil
	return result, err
}

func (userUseCase *UserUseCase) DeleteOneById(ctx context.Context, id string) (result *response.Response[*entity.User], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	deletedUser, deletedUserErr := userUseCase.UserRepository.DeleteOneById(transaction.Tx, id)
	if deletedUserErr != nil {
		transaction.TxErr = deletedUserErr
		result = nil
		err = deletedUserErr
		return result, err
	}
	if deletedUser == nil {
		rollbackErr := transaction.Tx.Rollback()
		if rollbackErr != nil {
			transaction.TxErr = rollbackErr
			result = nil
			err = rollbackErr
			return result, err
		}
		result = &response.Response[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase DeleteOneById is failed, user is not deleted by id.",
			Data:    nil,
		}
		err = nil
		return result, err
	}

	result = &response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase DeleteOneById is succeed.",
		Data:    deletedUser,
	}
	err = nil
	return result, err
}
