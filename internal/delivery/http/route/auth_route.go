package route

import (
	"social-media/internal/delivery/http"

	"github.com/gorilla/mux"
)

type AuthRoute struct {
	Router         *mux.Router
	AuthController *http.AuthController
}

func NewAuthRoute(router *mux.Router, AuthController *http.AuthController) *AuthRoute {
	AuthRoute := &AuthRoute{
		Router:         router.PathPrefix("/auths").Subrouter(),
		AuthController: AuthController,
	}
	return AuthRoute
}

func (AuthRoute *AuthRoute) Register() {
	AuthRoute.Router.HandleFunc("/register", AuthRoute.AuthController.Register).Methods("POST")
	AuthRoute.Router.HandleFunc("/login", AuthRoute.AuthController.Login).Methods("POST")
}
