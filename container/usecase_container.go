package container

import "social-media/internal/use_case"

type UseCaseContainer struct {
	User *use_case.UserUseCase
	Auth *use_case.AuthUseCase
	Post *use_case.PostUseCase
}

func NewUseCaseContainer(
	user *use_case.UserUseCase,
	auth *use_case.AuthUseCase,
	post *use_case.PostUseCase,
) *UseCaseContainer {
	useCaseContainer := &UseCaseContainer{
		User: user,
		Auth: auth,
		Post: post,
	}
	return useCaseContainer
}
