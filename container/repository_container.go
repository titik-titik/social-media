package container

import "social-media/internal/repository"

type RepositoryContainer struct {
	User   *repository.UserRepository
	Search *repository.SearchRepository
}

func NewRepositoryContainer(
	user *repository.UserRepository,
	search *repository.SearchRepository,
) *RepositoryContainer {
	return &RepositoryContainer{
		User:   user,
		Search: search,
	}
}
