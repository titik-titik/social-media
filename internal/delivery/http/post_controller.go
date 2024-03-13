package http

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
	model_request "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/guregu/null"
)

type PostController struct {
	PostUseCase *use_case.PostUseCase
	Log         *zerolog.Logger
}

func NewPostController(useCase *use_case.PostUseCase, log *zerolog.Logger) *PostController {
	return &PostController{
		PostUseCase: useCase,
		Log:         log,
	}
}

func (p *PostController) Create(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.CreatePostRequest)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.Log.Warn().Msgf("failed to parse body request : %+v", err)
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	ctx := r.Context()
	createdPost, createdPostErr := p.PostUseCase.Create(ctx, req)
	if createdPostErr == nil {
		response.NewResponse(w, createdPost)
	}
}

func (p *PostController) Find(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.GetPostRequest)
	postId := mux.Vars(r)["id"]

	if postId == "" {
		p.Log.Warn().Msgf("failed to parse param request : %+v", postId)
		return
	}

	req.PostId = null.NewString(postId, true)

	ctx := r.Context()
	foundPost, foundPostErr := p.PostUseCase.Find(ctx, req)
	if foundPostErr == nil {
		response.NewResponse(w, foundPost)
	}
}

func (p *PostController) Get(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.GetAllPostRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.Log.Warn().Msgf("failed to parse body request : %+v", err)
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	ctx := r.Context()
	foundPosts, foundPostsErr := p.PostUseCase.Get(ctx, req)
	if foundPostsErr == nil {
		response.NewResponse(w, foundPosts)
	}
}

func (p PostController) Update(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.UpdatePostRequest)
	postId := mux.Vars(r)["id"]

	if postId == "" {
		p.Log.Warn().Msgf("failed to parse param request : %+v", postId)
		return
	}

	req.ID = postId

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		p.Log.Warn().Msgf("failed to parse body request : %+v", err)
		response.NewResponse(w, &response.Response[*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		})
		return
	}

	ctx := r.Context()
	updatedPost, updatedPostErr := p.PostUseCase.Update(ctx, req)
	if updatedPostErr == nil {
		response.NewResponse(w, updatedPost)
	}
}

func (p PostController) Delete(w http.ResponseWriter, r *http.Request) {
	req := new(model_request.DeletePostRequest)
	postId := mux.Vars(r)["id"]

	if postId == "" {
		p.Log.Warn().Msgf("failed to parse param request : %+v", postId)
		return
	}

	req.ID = postId

	ctx := r.Context()
	deletedPost, deletedPostErr := p.PostUseCase.Delete(ctx, req)
	if deletedPostErr == nil {
		response.NewResponse(w, deletedPost)
	}
}
