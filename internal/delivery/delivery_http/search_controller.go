package delivery_http

import "social-media/internal/use_case"

type SearchController struct {
	SearchUseCase *use_case.SearchUseCase
}

func NewSearchController(searchUseCase *use_case.SearchUseCase) *SearchController {
	searchController := &SearchController{
		SearchUseCase: searchUseCase,
	}
	return searchController
}
