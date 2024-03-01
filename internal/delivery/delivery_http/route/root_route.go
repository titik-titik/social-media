package route

import (
	"github.com/gorilla/mux"
)

type RootRoute struct {
	Router    *mux.Router
	UserRoute *UserRoute
}

func NewRootRoute(
	router *mux.Router,
	userRoute *UserRoute,
) *RootRoute {
	rootRoute := &RootRoute{
		Router:    router,
		UserRoute: userRoute,
	}
	return rootRoute
}

func (rootRoute *RootRoute) Register() {
	rootRoute.UserRoute.Register()
}
