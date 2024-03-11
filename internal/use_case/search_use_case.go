package use_case

import (
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model/response"
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

func (searchUseCase *SearchUseCase) FindManyUser() (result *response.Response[[]*entity.User]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := searchUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result = nil
			return err
		}

		foundAllUser, err := searchUseCase.SearchRepository.FindManyUser(begin)
		if err != nil {
			return err
		}
		if foundAllUser == nil {
			result = &response.Response[[]*entity.User]{
				Code:    http.StatusNotFound,
				Message: "SearchUserCase FindManyUser is failed, user is not found.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &response.Response[[]*entity.User]{
			Code:    http.StatusOK,
			Message: "SearchUserCase FindManyUser is succeed.",
			Data:    foundAllUser,
		}
		return err
	})

	if beginErr != nil {
		result = &response.Response[[]*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "SearchUserCase FindManyUser  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}

func (searchUseCase *SearchUseCase) FindManyPostByUserId(id string) (result *response.Response[[]*entity.Post]) {
	beginErr := crdb.Execute(func() (err error) {
		begin, err := searchUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
		if err != nil {
			result = nil
			return err
		}

		foundAllPost, err := searchUseCase.SearchRepository.FindManyPostByUserId(begin, id)
		if err != nil {
			return err
		}
		if foundAllPost == nil {
			result = &response.Response[[]*entity.Post]{
				Code:    http.StatusNotFound,
				Message: "SearchPostCase FindManyPostByUserId is failed, post is not found by user id.",
				Data:    nil,
			}
			return err
		}

		err = begin.Commit()
		result = &response.Response[[]*entity.Post]{
			Code:    http.StatusOK,
			Message: "SearchPostCase FindManyPostByUserId is succeed",
			Data:    foundAllPost,
		}
		return err
	})

	if beginErr != nil {
		result = &response.Response[[]*entity.Post]{
			Code:    http.StatusInternalServerError,
			Message: "SearchPostCase FindManyPostByUserId  is failed, " + beginErr.Error(),
			Data:    nil,
		}
	}

	return result
}
