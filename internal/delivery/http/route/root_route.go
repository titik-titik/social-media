package route

import (
	"github.com/gorilla/mux"
)

type RootRoute struct {
	Router    *mux.Router
	UserRoute *UserRoute
	AuthRoute *AuthRoute
	PostRoute *PostRoute
}

func NewRootRoute(
	router *mux.Router,
	userRoute *UserRoute,
	authRoute *AuthRoute,
	postRoute *PostRoute,
) *RootRoute {
	rootRoute := &RootRoute{
		Router:    router,
		UserRoute: userRoute,
		AuthRoute: authRoute,
		PostRoute: postRoute,
	}
	return rootRoute
}

func (rootRoute *RootRoute) Register() {
	rootRoute.UserRoute.Register()
	rootRoute.AuthRoute.Register()
	// rootRoute.PostRoute.Register()
}
