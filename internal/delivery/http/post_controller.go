package http

import (
	"encoding/json"
	"net/http"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/guregu/null"
)

type PostController struct {
	PostUseCase *use_case.PostUseCase
}

func NewPostController(useCase *use_case.PostUseCase) *PostController {
	return &PostController{
		PostUseCase: useCase,
	}
}

func (p *PostController) Create(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.CreatePostRequest)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	result := p.PostUseCase.Create(req)

	response.NewResponse(w, result)
}

func (p *PostController) Find(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.GetPostRequest)
	postId := mux.Vars(r)["id"]

	req.PostId = null.NewString(postId, true)

	result := p.PostUseCase.Find(req)

	response.NewResponse(w, result)
}

func (p *PostController) Get(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.GetAllPostRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	result := p.PostUseCase.Get(req)

	response.NewResponse(w, result)
}

func (p PostController) Update(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.UpdatePostRequest)
	postId := mux.Vars(r)["id"]

	req.ID = postId

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	result := p.PostUseCase.Update(req)

	response.NewResponse(w, result)
}

func (p PostController) Delete(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.DeletePostRequest)
	postId := mux.Vars(r)["id"]

	req.ID = postId

	if err := p.PostUseCase.Delete(req); err != nil {
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	response.NewResponse(w, &response.Response[response.PostResponse]{
		Message: http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
	})
}
