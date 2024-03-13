package route

import (
	"github.com/gorilla/mux"
	"social-media/internal/delivery/http"
)

type UserRoute struct {
	Router         *mux.Router
	UserController *http.UserController
}

func NewUserRoute(router *mux.Router, userController *http.UserController) *UserRoute {
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
