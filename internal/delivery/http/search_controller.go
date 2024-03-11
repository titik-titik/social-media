package http

import (
	"social-media/internal/use_case"
)

type SearchController struct {
	UserUseCase *use_case.UserUseCase
}

func NewSearchController(sc *use_case.searchUseCase) *SearchController {
	SearchController := &SearchController{
		UserUseCase: sc,
	}
	return SearchController
}
