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
		Router:         router.PathPrefix("/auth").Subrouter(),
		AuthController: AuthController,
	}
	return AuthRoute
}

func (AuthRoute *AuthRoute) Register() {
}
