package use_case

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
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
	Log            *zerolog.Logger
	Validate       *validator.Validate
}

func NewPostUseCase(db *config.DatabaseConfig, postRepository *repository.PostRepository, log *zerolog.Logger, validate *validator.Validate) *PostUseCase {
	return &PostUseCase{
		PostRepository: postRepository,
		DB:             db,
		Log:            log,
		Validate:       validate,
	}
}

func (p *PostUseCase) Create(request *model_controller.CreatePostRequest) *response.Response[*response.PostResponse] {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	post := &entity.Post{
		Id:          null.StringFrom(uuid.NewString()),
		UserId:      null.StringFrom("c9bce534-36a8-43c1-b7d3-071e86673074"),
		ImageUrl:    request.ImageUrl,
		Description: request.Description,
		CreatedAt:   null.NewTime(time.Now(), true),
		UpdatedAt:   null.NewTime(time.Now(), true),
	}

	if err = p.PostRepository.Create(tx, post); err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to create new post : %+v, unable to back : %+v", err, rollbackErr)
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := tx.Commit(); err != nil {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}

func (p *PostUseCase) Find(request *model_controller.GetPostRequest) *response.Response[*response.PostResponse] {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		panic(err)
	}

	post := &entity.Post{}

	if err = p.PostRepository.FindByID(tx, post, request.PostId); err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to find by id post : %+v, unable to back : %+v", err, rollbackErr)
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    converter.PostToResponse(post),
	}
}

func (p PostUseCase) Get(request *model_controller.GetAllPostRequest) *response.Response[[]*response.PostResponse] {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return &response.Response[[]*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := p.Validate.Struct(request); err != nil {
		p.Log.Error().Err(err).Msgf("failed to validate request body")
		return &response.Response[[]*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  err.Error(),
		}
	}

	posts := new([]entity.Post)

	if err = p.PostRepository.Get(tx, posts, request.Order, request.Limit, request.Offset); err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to get all post : %+v, unable to rollback : %+v", err, rollbackErr)
		return &response.Response[[]*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := tx.Commit(); err != nil {
		return &response.Response[[]*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return &response.Response[[]*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    converter.PostToResponses(posts),
	}
}

func (p PostUseCase) Update(request *model_controller.UpdatePostRequest) *response.Response[*response.PostResponse] {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	total, err := p.PostRepository.CountByID(tx, request.ID)

	if err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to count by id post : %+v, unable to rollback : %+v", err, rollbackErr)
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if total < 0 {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	}

	post := &entity.Post{
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
		UpdatedAt:   null.NewTime(time.Now(), true),
	}

	if err = p.PostRepository.Update(tx, post, request.ID); err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to update post : %+v, unable to rollback : %+v", err, rollbackErr)
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := tx.Commit(); err != nil {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}

func (p PostUseCase) Delete(request *model_controller.DeletePostRequest) *response.Response[*response.PostResponse] {
	tx, err := p.DB.CockroachdbDatabase.Connection.Begin()

	if err != nil {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	total, err := p.PostRepository.CountByID(tx, request.ID)

	if err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to count by id post : %+v, unable to rollback : %+v", err, rollbackErr)
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if total < 0 {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
	}

	if err = p.PostRepository.Delete(tx, request.ID); err != nil {
		rollbackErr := tx.Rollback()
		p.Log.Error().Msgf("failed to delete post by id : %+v,unable to rollback : %+v", err, rollbackErr)
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if err := tx.Commit(); err != nil {
		return &response.Response[*response.PostResponse]{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	return &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
}
