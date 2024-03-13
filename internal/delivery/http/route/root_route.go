package route

import (
	"github.com/gorilla/mux"
	"social-media/internal/delivery/http/middleware"
)

type RootRoute struct {
	Router           *mux.Router
	RootMiddleware   *middleware.RootMiddleware
	ProtectedRoute   *ProtectedRoute
	UnprotectedRoute *UnprotectedRoute
}

func NewRootRoute(
	router *mux.Router,
	rootMiddleware *middleware.RootMiddleware,
	protectedRoute *ProtectedRoute,
	unprotectedRoute *UnprotectedRoute,
) *RootRoute {
	rootRoute := &RootRoute{
		Router:           router,
		RootMiddleware:   rootMiddleware,
		ProtectedRoute:   protectedRoute,
		UnprotectedRoute: unprotectedRoute,
	}
	return rootRoute
}

func (rootRoute *RootRoute) Register() {
	rootRoute.Router.Use(rootRoute.RootMiddleware.TransactionMiddleware.GetMiddleware)
	rootRoute.UnprotectedRoute.Register()
	rootRoute.ProtectedRoute.Router.Use(rootRoute.RootMiddleware.AuthMiddleware.GetMiddleware)
	rootRoute.ProtectedRoute.Register()
}
