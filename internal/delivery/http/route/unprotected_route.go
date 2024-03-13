package route

import "github.com/gorilla/mux"

type UnprotectedRoute struct {
	Router    *mux.Router
	AuthRoute *AuthRoute
}

func NewUnprotectedRoute(router *mux.Router, authRoute *AuthRoute) *UnprotectedRoute {
	protectedRoute := &UnprotectedRoute{
		Router:    router.Path("").Subrouter(),
		AuthRoute: authRoute,
	}
	return protectedRoute
}

func (protectedRoute *UnprotectedRoute) Register() {
	protectedRoute.AuthRoute.Register()
}
