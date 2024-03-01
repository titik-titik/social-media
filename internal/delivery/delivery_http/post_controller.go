package delivery_http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/guregu/null"
	"net/http"
	"social-media/internal/model/request"
	"social-media/internal/model/response"
	"social-media/internal/use_case"
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
	var req request.CreatePostRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := p.PostUseCase.Create(&req); err != nil {
		http.Error(w, "Failed to create new post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data := response.NewResponse(http.StatusText(http.StatusCreated), new(string))

	res, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Failed to create new post", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)

	if err != nil {
		http.Error(w, "Failed to create new post", http.StatusInternalServerError)
		return
	}
}

func (p *PostController) Get(w http.ResponseWriter, r *http.Request) {
	var req request.GetPostRequest
	postId := mux.Vars(r)["id"]

	req.PostId = null.NewString(postId, true)

	post, errGet := p.PostUseCase.Get(&req)

	if errGet != nil {
		http.Error(w, "Failed to get post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := response.NewResponse(http.StatusText(http.StatusOK), post)

	res, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "Failed to get post", http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)

	if err != nil {
		http.Error(w, "Failed to get post", http.StatusInternalServerError)
		return
	}
}
