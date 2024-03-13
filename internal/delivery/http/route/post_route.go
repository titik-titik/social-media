package route

import (
	"github.com/gorilla/mux"
	"social-media/internal/delivery/http"
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
	postRoute.Router.HandleFunc("/", postRoute.PostController.Get).Methods("GET")
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Find).Methods("GET")
	postRoute.Router.HandleFunc("/", postRoute.PostController.Create).Methods("POST")
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Update).Methods("PUT")
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Delete).Methods("DELETE")
}
