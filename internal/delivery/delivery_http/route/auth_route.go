package route

//import (
//	"github.com/gorilla/mux"
//	"social-media/internal/delivery/delivery_http"
//)
//
//type AuthRoute struct {
//	Router         *mux.Router
//	AuthController *delivery_http.AuthController
//}
//
//func NewAuthRoute(router *mux.Router, AuthController *delivery_http.AuthController) *AuthRoute {
//	AuthRoute := &AuthRoute{
//		Router:         router.PathPrefix("/auth").Subrouter(),
//		AuthController: AuthController,
//	}
//	return AuthRoute
//}
//
//func (AuthRoute *AuthRoute) Register() {
//	AuthRoute.Router.HandleFunc("/register", AuthRoute.AuthController.Register).Methods("POST")
//}
