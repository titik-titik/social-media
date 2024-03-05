package container

import "social-media/internal/delivery/http"

type ControllerContainer struct {
	User *http.UserController
	Auth *http.AuthController
}

func NewControllerContainer(
	user *http.UserController,
	auth *http.AuthController,
) *ControllerContainer {
	controllerContainer := &ControllerContainer{
		User: user,
		Auth: auth,
	}
	return controllerContainer
}
