package repository

import (
	"social-media/internal/entity"
)

type LoginRepositoryRequest struct {
	Session *entity.Session
	User    *entity.User
}
