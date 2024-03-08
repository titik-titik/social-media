package use_case

import (
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

func (searchUseCase *SearchUseCase) FindAllUser() *model.Result[[]*entity.User] {
	begin, beginErr := searchUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[[]*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "SearchUserCase FindAllUser is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundAllUser := searchUseCase.SearchRepository.FindAllUser(begin)

	if foundAllUser == nil {
		return &model.Result[[]*entity.User]{
			Code:    http.StatusNotFound,
			Message: "SearchUserCase FindAllUser is failed",
			Data:    nil,
		}
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		return &model.Result[[]*entity.User]{
			Code:    http.StatusInternalServerError,
			Message: "SearchUserCase FindAllUser is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[[]*entity.User]{
		Code:    http.StatusOK,
		Message: "SearchUserCase FindAllUser is succeed.",
		Data:    foundAllUser,
	}
}

func (searchUseCase *SearchUseCase) FindAllPostByUserId(id string) *model.Result[[]*entity.Post] {
	begin, beginErr := searchUseCase.DatabaseConfig.CockroachdbDatabase.Connection.Begin()
	if beginErr != nil {
		return &model.Result[[]*entity.Post]{
			Code:    http.StatusInternalServerError,
			Message: "SearchPostCase FindAllPostByUserId is failed, transaction begin is failed.",
			Data:    nil,
		}
	}

	foundAllPost := searchUseCase.SearchRepository.FindAllPostByUserId(begin, id)

	if foundAllPost == nil {
		return &model.Result[[]*entity.Post]{
			Code:    http.StatusNotFound,
			Message: "SearchPostCase FindAllPostByUserId is failed",
			Data:    nil,
		}
	}

	commitErr := begin.Commit()
	if commitErr != nil {
		return &model.Result[[]*entity.Post]{
			Code:    http.StatusInternalServerError,
			Message: "SearchPostCase FindAllPostByUserId is failed, transaction commit is failed.",
			Data:    nil,
		}
	}

	return &model.Result[[]*entity.Post]{
		Code:    http.StatusOK,
		Message: "SearchPostCase FindAllPostByUserId is succeed",
		Data:    foundAllPost,
	}
}
