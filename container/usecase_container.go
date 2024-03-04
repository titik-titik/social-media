package container

import "social-media/internal/use_case"

type UseCaseContainer struct {
	User   *use_case.UserUseCase
	Search *use_case.SearchUseCase
}

func NewUseCaseContainer(
	user *use_case.UserUseCase,
	search *use_case.SearchUseCase,
) *UseCaseContainer {
	return &UseCaseContainer{
		User:   user,
		Search: search,
	}
}
