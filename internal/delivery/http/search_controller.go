package http

import (
	"net/http"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"

	"github.com/gorilla/mux"
)

type SearchController struct {
	searchUseCase *use_case.
}

func NewSearchController(sc *use_case.searchUseCase) *SearchController {
	SearchController := &SearchController{
		searchUseCase: sc,
	}
	return SearchController
}
func (SearchController *SearchController) Get(writer http.ResponseWriter, reader *http.Request) {
	limitUser := mux.Vars(reader)["limit"]
	offsetUser := mux.Vars(reader)["offset"]
	orderUser := mux.Vars(reader)["order"]

	request := model_request.GetAllUserRequest{
		Limit:  limitUser,
		Offset: offsetUser,
		Order:  orderUser,
	}
	result := SearchController.searchUseCase.find(request)
	response.NewResponse(writer, result)

	response.NewResponse(writer, result)
}
