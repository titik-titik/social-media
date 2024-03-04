package container

import "social-media/internal/delivery/delivery_http"

type ControllerContainer struct {
	User *delivery_http.UserController
}

func NewControllerContainer(
	user *delivery_http.UserController,
) *ControllerContainer {
	return &ControllerContainer{
		User: user,
	}
}
