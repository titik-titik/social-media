package container

import "social-media/internal/repository"

type RepositoryContainer struct {
	User    *repository.UserRepository
	Session *repository.SessionRepository
	Post    *repository.PostRepository
}

func NewRepositoryContainer(
	user *repository.UserRepository,
	session *repository.SessionRepository,
	post *repository.PostRepository,
) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		User:    user,
		Session: session,
		Post:    post,
	}
	return repositoryContainer
}
