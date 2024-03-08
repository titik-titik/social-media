package use_case

import (
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/repository"
	"time"

	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
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

func (userUseCase *UserUseCase) FindOneById(id string) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneById is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneById(begin, id)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase FindOneById is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneById is failed, user is not found by id.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneById is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneById is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByUsername(username string) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByUsername is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneByUsername(begin, username)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase FindOneByUsername is failed, transaction rollback is failed.",
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByUsername is failed, user is not found by username.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByUsername is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByUsername is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByEmail(email string) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByEmail is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneByEmail(begin, email)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase FindOneByEmail is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByEmail is failed, user is not found by email.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByEmail is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByEmail is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByEmailAndPassword(email, password string) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByEmailAndPassword is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneByEmailAndPassword(begin, email, password)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase FindOneByEmailAndPassword is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByEmailAndPassword is failed, user is not found by email and password.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByEmailAndPassword is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByEmailAndPassword is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) FindOneByUsernameAndPassword(username, password string) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByUsernameAndPassword is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneByUsernameAndPassword(begin, username, password)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase FindOneByUsernameAndPassword is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase FindOneByUsernameAndPassword is failed, user is not found by username and password.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase FindOneByUsernameAndPassword is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase FindOneByUsernameAndPassword is succeed.",
		Data:    foundUser,
	}
}

func (userUseCase *UserUseCase) CreateOne(toCreateUser *entity.User) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase CreateOne is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	createdUser := userUseCase.UserRepository.CreateOne(begin, toCreateUser)
	if createdUser == nil {
		rollbackErr := begin.Rollback()
		if rollbackErr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase CreateOne is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase CreateOne is failed, user is not created.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase CreateOne is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase CreateOne is succeed.",
		Data:    createdUser,
	}
}

func (userUseCase *UserUseCase) PatchOneById(id string, toPatchUser *entity.User) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneById is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	patchedUser := userUseCase.UserRepository.PatchOneById(begin, id, toPatchUser)
	if patchedUser == nil {
		rollbackErr := begin.Rollback()
		if rollbackErr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase PatchOneById is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase UserPatchOneByIdRequest is failed, user is not patched.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneById is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase UserPatchOneByIdRequest is succeed.",
		Data:    patchedUser,
	}
}

func (userUseCase *UserUseCase) PatchOneByIdFromRequest(id string, request *model_request.UserPatchOneByIdRequest) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneById(begin, id)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase PatchOneByIdFromRequest is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, user is not found by id.",
			Data:    nil,
		}
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
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, password hashing is failed.",
			Data:    nil,
		}
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
		rollbackErr := begin.Rollback()
		if rollbackErr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase PatchOneByIdFromRequest is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, user is not patched.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase PatchOneByIdFromRequest is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase PatchOneByIdFromRequest is succeed.",
		Data:    patchedUser,
	}
}

func (userUseCase *UserUseCase) DeleteOneById(id string) *model.Result[*entity.User] {
	begin, beginErr := userUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase DeleteOneById is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundUser := userUseCase.UserRepository.FindOneById(begin, id)
	if foundUser == nil {
		rollbackEr := begin.Rollback()
		if rollbackEr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase DeleteOneById is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusNotFound,
			Message: "UserUserCase DeleteOneById is failed, user is not found by id.",
			Data:    nil,
		}
	}

	deletedUser := userUseCase.UserRepository.DeleteOneById(begin, id)
	if deletedUser == nil {
		rollbackErr := begin.Rollback()
		if rollbackErr != nil {
			return &model.Result[*entity.User]{
				Code:    http.StatusInternalServerError,
				Message: "UserUserCase DeleteOneById is failed, transaction rollback is failed.",
				Data:    nil,
			}
		}
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase DeleteOneById is failed, user is not deleted.",
			Data:    nil,
		}
	}

	commitErr := crdb.Execute(func() (err error) {
		err = begin.Commit()
		return err
	})
	if commitErr != nil {
		return &model.Result[*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "UserUserCase DeleteOneById is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[*entity.User]{
		Code:    http.StatusOK,
		Message: "UserUserCase DeleteOneById is succeed.",
		Data:    deletedUser,
	}
}
