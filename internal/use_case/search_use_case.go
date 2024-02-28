package use_case

import (
	"social-media/internal/entity"
	"social-media/internal/model"
	"social-media/internal/repository"
)

type SearchUseCase struct {
	SearchRepository *repository.SearchRepository
}

func NewSearchUseCase(
	searchRepository *repository.SearchRepository,
) *SearchUseCase {
	searchUseCase := &SearchUseCase{
		SearchRepository: searchRepository,
	}
	return searchUseCase
}

func (searchUseCase *SearchUseCase) FindAllUser() *model.Result[*[]entity.User] {
	foundAllUser := searchUseCase.SearchRepository.FindAllUser()

	if foundAllUser == nil {
		return &model.Result[*[]entity.User]{
			Code:    404,
			Message: "SearchUserCase FindAllUser is failed",
			Data:    nil,
		}
	}

	return &model.Result[*[]entity.User]{
		Code:    200,
		Message: "SearchUserCase FindAllUser is succeed.",
		Data:    foundAllUser,
	}
}

func (searchUseCase *SearchUseCase) FindAllPostByUserId(id string) *model.Result[*[]entity.Post] {
	foundAllPost := searchUseCase.SearchRepository.FindAllPostByUserId(id)

	if foundAllPost == nil {
		return &model.Result[*[]entity.Post]{
			Code: 404,
			Message: "SearchPostCase FindAllPostByUserId is failed",
			Data: nil,
		}
	}

	return &model.Result[*[]entity.Post]{
		Code: 200,
		Message: "SearchPostCase FindAllPostByUserId is succeed",
		Data: foundAllPost,
	}
}