package container

import "social-media/internal/repository"

type RepositoryContainer struct {
	User    *repository.UserRepository
	Session *repository.SessionRepository
	Post    *repository.PostRepository
	Search  *repository.SearchRepository
}

func NewRepositoryContainer(
	user *repository.UserRepository,
	session *repository.SessionRepository,
	post *repository.PostRepository,
	search *repository.SearchRepository,
) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		User:    user,
		Session: session,
		Post:    post,
		Search:  search,
	}
	return repositoryContainer
}
