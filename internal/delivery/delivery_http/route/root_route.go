package route

import (
	"github.com/gorilla/mux"
)

type RootRoute struct {
	Router    *mux.Router
	UserRoute *UserRoute
	PostRoute *PostRoute
}

func NewRootRoute(
	router *mux.Router,
	userRoute *UserRoute,
	postRoute *PostRoute,
) *RootRoute {
	rootRoute := &RootRoute{
		Router:    router,
		UserRoute: userRoute,
		PostRoute: postRoute,
	}
	return rootRoute
}

func (rootRoute *RootRoute) Register() {
	rootRoute.UserRoute.Register()
	rootRoute.PostRoute.Register()
}
