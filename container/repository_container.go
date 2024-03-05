package container

import "social-media/internal/repository"

type RepositoryContainer struct {
	User   *repository.UserRepository
	Auth   *repository.AuthRepository
	Search *repository.SearchRepository
}

func NewRepositoryContainer(
	user *repository.UserRepository,
	auth *repository.AuthRepository,
	search *repository.SearchRepository,
) *RepositoryContainer {
	repositoryContainer := &RepositoryContainer{
		User:   user,
		Auth:   auth,
		Search: search,
	}
	return repositoryContainer
}
