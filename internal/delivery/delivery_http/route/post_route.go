package route

import (
	"social-media/internal/delivery/delivery_http"

	"github.com/gorilla/mux"
)

type PostRoute struct {
	Router         *mux.Router
	PostController *delivery_http.PostController
}

func NewPostRoute(router *mux.Router, userController *delivery_http.PostController) *PostRoute {
	userRoute := &PostRoute{
		Router:         router.PathPrefix("/posts").Subrouter(),
		PostController: userController,
	}
	return userRoute
}

func (userRoute *PostRoute) Register() {
	userRoute.Router.HandleFunc("/{id}", userRoute.PostController.Get).Methods("GET")
	userRoute.Router.HandleFunc("/", userRoute.PostController.Create).Methods("POST")

}
