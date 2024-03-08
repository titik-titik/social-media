package repository

import (
	"social-media/internal/entity"
	model_controller "social-media/internal/model/request/controller"
)

type LoginRepositoryRequest struct {
	LoginControllerRequest *model_controller.LoginRequest
	Session                *entity.Session
	User                   *entity.User
}
