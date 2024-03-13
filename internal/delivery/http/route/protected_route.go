package route

import "github.com/gorilla/mux"

type ProtectedRoute struct {
	Router    *mux.Router
	UserRoute *UserRoute
	PostRoute *PostRoute
}

func NewProtectedRoute(
	router *mux.Router,
	userRoute *UserRoute,
	postRoute *PostRoute,
) *ProtectedRoute {
	protectedRoute := &ProtectedRoute{
		Router:    router.Path("").Subrouter(),
		UserRoute: userRoute,
		PostRoute: postRoute,
	}
	return protectedRoute
}

func (protectedRoute *ProtectedRoute) Register() {
	protectedRoute.UserRoute.Register()
	protectedRoute.PostRoute.Register()
}
