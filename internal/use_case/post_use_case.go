package use_case

import "social-media/internal/repository"

type PostUseCase struct {
	PostRepository *repository.PostRepository
}

func NewPostUseCase(postRepository *repository.PostRepository) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
	}
}

func (p *PostUseCase) Create()
