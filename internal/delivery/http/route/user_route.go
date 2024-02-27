package route

import (
	http_delivery "social-media/internal/delivery/http"

	"github.com/gorilla/mux"
)

type UserRoute struct {
	Router         *mux.Router
	UserController *http_delivery.UserController
}

func NewUserRoute(router *mux.Router, userController *http_delivery.UserController) *UserRoute {
	userRoute := &UserRoute{
		Router:         router.PathPrefix("/users").Subrouter(),
		UserController: userController,
	}
	return userRoute
}

func (userRoute *UserRoute) Register() {
	userRoute.Router.HandleFunc("/{id}", userRoute.UserController.FindOneById).Methods("GET")
	userRoute.Router.HandleFunc("", userRoute.UserController.FindOneByOneParam).Methods("GET")
	userRoute.Router.HandleFunc("/{id}", userRoute.UserController.PatchOneById).Methods("PATCH")
	userRoute.Router.HandleFunc("/{id}", userRoute.UserController.DeleteOneById).Methods("DELETE")
}
