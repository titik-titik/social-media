package use_case

import (
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/repository"
	"time"
)

type UserUseCase struct {
	DatabaseConfig *config.DatabaseConfig
	UserRepository *repository.UserRepository
}

func NewUserUseCase(
	databaseConfig *config.DatabaseConfig,
	userRepository *repository.UserRepository,
) *UserUseCase {
	userUseCase := &UserUseCase{
		DatabaseConfig: databaseConfig,
		UserRepository: userRepository,
	}
	return userUseCase
}

func (userUseCase *UserUseCase) FindOneById(id string) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		foundUser := userUseCase.UserRepository.FindOneById(begin, id)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase FindOneById is failed, user is not found by id.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase FindOneById is succeed.",
			Data:    foundUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneById is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) FindOneByUsername(username string) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		foundUser := userUseCase.UserRepository.FindOneByUsername(begin, username)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase FindOneByUsername is failed, user is not found by username.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase FindOneByUsername is succeed.",
			Data:    foundUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByUsername is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) FindOneByEmail(email string) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		foundUser := userUseCase.UserRepository.FindOneByEmail(begin, email)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase FindOneByEmail is failed, user is not found by email.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase FindOneByEmail is succeed.",
			Data:    foundUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByEmail is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) FindOneByEmailAndPassword(email, password string) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		foundUser := userUseCase.UserRepository.FindOneByEmailAndPassword(begin, email, password)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase FindOneByEmailAndPassword is failed, user is not found by email and password.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase FindOneByEmailAndPassword is succeed.",
			Data:    foundUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByEmailAndPassword is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) FindOneByUsernameAndPassword(username, password string) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		foundUser := userUseCase.UserRepository.FindOneByUsernameAndPassword(begin, username, password)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase FindOneByUsernameAndPassword is failed, user is not found by username and password.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase FindOneByUsernameAndPassword is succeed.",
			Data:    foundUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByUsernameAndPassword is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) CreateOne(toCreateUser *entity.User) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		createdUser := userUseCase.UserRepository.CreateOne(begin, toCreateUser)
		if createdUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase CreateOne is failed, user is not created.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase CreateOne is succeed.",
			Data:    createdUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase CreateOne is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) PatchOneById(id string, toPatchUser *entity.User) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		patchedUser := userUseCase.UserRepository.PatchOneById(begin, id, toPatchUser)
		if patchedUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase UserPatchOneByIdRequest is failed, user is not patched.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase UserPatchOneByIdRequest is succeed.",
			Data:    patchedUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneById is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) PatchOneByIdFromRequest(id string, request *model_request.UserPatchOneByIdRequest) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		foundUser := userUseCase.UserRepository.FindOneById(begin, id)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase PatchOneByIdFromRequest is failed, user is not found by id.",
				Data:    nil,
			}
			return err
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
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase PatchOneByIdFromRequest is failed, password hashing is failed.",
				Data:    nil,
			}
			return err
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

		foundUser.UpdatedAt = null.NewTime(time.Now().UTC(), true)

		patchedUser := userUseCase.UserRepository.PatchOneById(begin, id, foundUser)
		if patchedUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase PatchOneByIdFromRequest is failed, user is not patched.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase PatchOneByIdFromRequest is succeed.",
			Data:    patchedUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}

func (userUseCase *UserUseCase) DeleteOneById(id string) (result *model.Result[*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()

		foundUser := userUseCase.UserRepository.FindOneById(begin, id)
		if foundUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusNotFound,
				Message: "UserUserCase DeleteOneById is failed, user is not found by id.",
				Data:    nil,
			}
			return err
		}

		deletedUser := userUseCase.UserRepository.DeleteOneById(begin, id)
		if deletedUser == nil {
			err = begin.Rollback()
			result = &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase DeleteOneById is failed, user is not deleted.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[*entity.User]{
			Code:    http.StatusOK,
			Message: "UserUserCase DeleteOneById is succeed.",
			Data:    deletedUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase DeleteOneById is failed, transaction is failed.",
			Data:    nil,
		}
	}

	return result
}
