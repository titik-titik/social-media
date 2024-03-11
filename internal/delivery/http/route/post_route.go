package route

import (
	"social-media/internal/delivery/http"
	"social-media/internal/delivery/http/middleware"

	"github.com/gorilla/mux"
)

type PostRoute struct {
	Router         *mux.Router
	PostController *http.PostController
	Middleware     *middleware.AuthMiddleware
}

func NewPostRoute(router *mux.Router, postController *http.PostController, middleware *middleware.AuthMiddleware) *PostRoute {
	postRoute := &PostRoute{
		Router:         router.PathPrefix("/posts").Subrouter(),
		PostController: postController,
		Middleware:     middleware,
	}
	return postRoute
}

func (postRoute *PostRoute) Register() {
	postRoute.Router.Use(postRoute.Middleware.Middleware)
	postRoute.Router.HandleFunc("/", postRoute.PostController.Get).Methods("GET")
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Find).Methods("GET")
	postRoute.Router.HandleFunc("/", postRoute.PostController.Create).Methods("POST")
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Update).Methods("PUT")
	postRoute.Router.HandleFunc("/{id}", postRoute.PostController.Delete).Methods("DELETE")
}
