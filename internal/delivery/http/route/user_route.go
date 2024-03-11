package route

import (
	"social-media/internal/delivery/http"
	"social-media/internal/delivery/http/middleware"

	"github.com/gorilla/mux"
)

type UserRoute struct {
	Router         *mux.Router
	UserController *http.UserController
	Middleware     *middleware.AuthMiddleware
}

func NewUserRoute(router *mux.Router, userController *http.UserController, middleware *middleware.AuthMiddleware) *UserRoute {
	userRoute := &UserRoute{
		Router:         router.PathPrefix("/users").Subrouter(),
		UserController: userController,
		Middleware:     middleware,
	}
	return userRoute
}

func (userRoute *UserRoute) Register() {
	userRoute.Router.Use(userRoute.Middleware.Middleware)
	userRoute.Router.HandleFunc("/{id}", userRoute.UserController.FindOneById).Methods("GET")
	userRoute.Router.HandleFunc("", userRoute.UserController.FindOneByOneParam).Methods("GET")
	userRoute.Router.HandleFunc("/{id}", userRoute.UserController.PatchOneById).Methods("PATCH")
	userRoute.Router.HandleFunc("/{id}", userRoute.UserController.DeleteOneById).Methods("DELETE")
}
