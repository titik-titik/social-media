package container

import "social-media/internal/repository"

type RepositoryContainer struct {
	User   *repository.UserRepository
	Post   *repository.PostRepository
	Auth   *repository.AuthRepository
	Search *repository.SearchRepository
}

func NewRepositoryContainer(
	user *repository.UserRepository,
	post *repository.PostRepository,
	auth *repository.AuthRepository,
	search *repository.SearchRepository,
) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		User:   user,
		Post:   post,
		Auth:   auth,
		Search: search,
	}
	return repositoryContainer
}
