package use_case

import (
	"errors"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model/converter"
	model_controller "social-media/internal/model/request/controller"
	"social-media/internal/model/response"
	"social-media/internal/repository"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
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

func (p *PostUseCase) Create(request *model_controller.CreatePostRequest) error {
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

func (p *PostUseCase) Find(request *model_controller.GetPostRequest) (*response.PostResponse, error) {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		panic(err)
	}

	post := &entity.Post{}

	if err = p.PostRepository.FindByID(tx, post, request.PostId); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return converter.PostToResponse(post), nil
}

func (p PostUseCase) Get(request *model_controller.GetAllPostRequest) ([]*response.PostResponse, error) {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	posts := new([]entity.Post)

	if err = p.PostRepository.Get(tx, posts, request.Order, request.Limit, request.Offset); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return converter.PostToResponses(posts), nil
}

func (p PostUseCase) Update(request *model_controller.UpdatePostRequest) error {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	total, err := p.PostRepository.CountByID(tx, request.ID)

	if err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	if total < 0 {
		return errors.New(http.StatusText(http.StatusNotFound))
	}

	post := &entity.Post{
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
		UpdatedAt:   null.NewTime(time.Now().UTC(), true),
	}

	if err = p.PostRepository.Update(tx, post, request.ID); err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	if err := tx.Commit(); err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}

func (p PostUseCase) Delete(request *model_controller.DeletePostRequest) error {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	total, err := p.PostRepository.CountByID(tx, request.ID)

	if err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	if total < 0 {
		return errors.New(http.StatusText(http.StatusNotFound))
	}

	if err = p.PostRepository.Delete(tx, request.ID); err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	if err := tx.Commit(); err != nil {
		return errors.New(http.StatusText(http.StatusInternalServerError))
	}

	return nil
}
