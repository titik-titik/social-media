package seeder

import (
	"social-media/internal/repository"
	"social-media/test/mock"
)

type PostSeeder struct {
	PostMock       *mock.PostMock
	PostRepository *repository.PostRepository
}

func NewPostSeeder(postRepository *repository.PostRepository) *PostSeeder {
	postSeeder := &PostSeeder{
		PostMock:       mock.NewPostMock(),
		PostRepository: postRepository,
	}
	return postSeeder
}

func (postSeeder *PostSeeder) Up() {
	//for _, post := range postSeeder.PostMock.Data {
	//	postSeeder.PostRepository.CreateOne(post)
	//}
}

func (postSeeder *PostSeeder) Down() {
	//for _, post := range postSeeder.PostMock.Data {
	//	postSeeder.PostRepository.DeleteOneById(post.Id.String)
	//}
}
