package route

import (
	"social-media/internal/delivery/http"

	"github.com/gorilla/mux"
)

type PostRoute struct {
	Router         *mux.Router
	PostController *http.PostController
}

func NewPostRoute(router *mux.Router, postController *http.PostController) *PostRoute {
	postRoute := &PostRoute{
		Router:         router.PathPrefix("/posts").Subrouter(),
		PostController: postController,
	}
	return postRoute
}

func (postRoute *PostRoute) Register() {
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Get).Methods("GET")
	postRoute.Router.HandleFunc("/", postRoute.PostController.Create).Methods("POST")

}
