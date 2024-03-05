package use_case

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/guregu/null"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model/converter"
	model_request "social-media/internal/model/request"
	"social-media/internal/model/response"
	"social-media/internal/repository"
	"time"
)

type PostUseCase struct {
	PostRepository *repository.PostRepository
	DB             *config.DatabaseConfig
}

func NewPostUseCase(db *config.DatabaseConfig, postRepository *repository.PostRepository) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
		DB:             db,
	}
}

func (p *PostUseCase) Create(request *model_request.CreatePostRequest) error {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		panic(err)
	}

	post := &entity.Post{
		Id:          null.StringFrom(uuid.NewString()),
		UserId:      null.StringFrom("c9bce534-36a8-43c1-b7d3-071e86673074"),
		ImageUrl:    request.ImageUrl,
		Description: request.Description,
		CreatedAt:   null.NewTime(time.Now().UTC(), true),
		UpdatedAt:   null.NewTime(time.Now().UTC(), true),
	}

	if err = p.PostRepository.Create(tx, post); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return nil
}

func (p *PostUseCase) Get(request *model_request.GetPostRequest) (*response.PostResponse, error) {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		panic(err)
	}

	post := &entity.Post{}

	if err = p.PostRepository.Get(tx, post, request.PostId); err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return converter.PostToResponse(post), nil
}
