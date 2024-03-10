package use_case

import (
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
	"social-media/internal/repository"
)

type SearchUseCase struct {
	DatabaseConfig   *config.DatabaseConfig
	SearchRepository *repository.SearchRepository
}

func NewSearchUseCase(
	databaseConfig *config.DatabaseConfig,
	searchRepository *repository.SearchRepository,
) *SearchUseCase {
	searchUseCase := &SearchUseCase{
		DatabaseConfig:   databaseConfig,
		SearchRepository: searchRepository,
	}
	return searchUseCase
}

func (searchUseCase *SearchUseCase) FindAllUser() (result *model.Result[[]*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := searchUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result = nil
			return err
		}

		foundAllUser, err := searchUseCase.SearchRepository.FindAllUser(begin)
		if err != nil {
			return err
		}
		if foundAllUser == nil {
			result = &model.Result[[]*entity.User]{
				Code:    http.StatusNotFound,
				Message: "SearchUserCase FindAllUser is failed, user is not found.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[[]*entity.User]{
			Code:    http.StatusOK,
			Message: "SearchUserCase FindAllUser is succeed.",
			Data:    foundAllUser,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[[]*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "SearchUserCase FindAllUser  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}

func (searchUseCase *SearchUseCase) FindAllPostByUserId(id string) (result *model.Result[[]*entity.Post]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := searchUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result = nil
			return err
		}

		foundAllPost, err := searchUseCase.SearchRepository.FindAllPostByUserId(begin, id)
		if err != nil {
			return err
		}
		if foundAllPost == nil {
			result = &model.Result[[]*entity.Post]{
				Code:    http.StatusNotFound,
				Message: "SearchPostCase FindAllPostByUserId is failed, post is not found by user id.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &model.Result[[]*entity.Post]{
			Code:    http.StatusOK,
			Message: "SearchPostCase FindAllPostByUserId is succeed",
			Data:    foundAllPost,
		}
		return err
	})

	if beginErr != nil {
		result = &model.Result[[]*entity.Post]{
			Code:    http.StatusInternalServerError,
			Message: "SearchPostCase FindAllPostByUserId  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}
