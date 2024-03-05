package container

import "social-media/internal/use_case"

type UseCaseContainer struct {
	User   *use_case.UserUseCase
	Auth   *use_case.AuthUseCase
	Search *use_case.SearchUseCase
}

func NewUseCaseContainer(
	user *use_case.UserUseCase,
	auth *use_case.AuthUseCase,
	search *use_case.SearchUseCase,
) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		User:   user,
		Auth:   auth,
		Search: search,
	}
	return useCaseContainer
}
