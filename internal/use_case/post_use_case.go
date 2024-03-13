package use_case

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"net/http"
	"social-media/internal/config"
	"social-media/internal/entity"
	"social-media/internal/model"
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

func (p *PostUseCase) Create(ctx context.Context, request *model_controller.CreatePostRequest) (result *response.Response[*response.PostResponse], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	post := &entity.Post{
		Id:          null.StringFrom(uuid.NewString()),
		UserId:      null.StringFrom("c9bce534-36a8-43c1-b7d3-071e86673074"),
		ImageUrl:    request.ImageUrl,
		Description: request.Description,
		CreatedAt:   null.NewTime(time.Now(), true),
		UpdatedAt:   null.NewTime(time.Now(), true),
	}

	createdPostErr := p.PostRepository.Create(transaction.Tx, post)
	if createdPostErr != nil {
		transaction.TxErr = createdPostErr
		result = nil
		err = createdPostErr
		return result, err
	}

	result = &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}
	err = nil
	return result, err
}

func (p *PostUseCase) Find(ctx context.Context, request *model_controller.GetPostRequest) (result *response.Response[*response.PostResponse], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	post := &entity.Post{}

	foundPostErr := p.PostRepository.FindByID(transaction.Tx, post, request.PostId)
	if foundPostErr != nil {
		transaction.TxErr = foundPostErr
		result = nil
		err = foundPostErr
		return result, err
	}

	result = &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    converter.PostToResponse(post),
	}
	err = nil
	return result, err
}

func (p PostUseCase) Get(ctx context.Context, request *model_controller.GetAllPostRequest) (result *response.Response[[]*response.PostResponse], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	if err := p.Validate.Struct(request); err != nil {
		p.Log.Error().Err(err).Msgf("failed to validate request body")
		result = &response.Response[[]*response.PostResponse]{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Errors:  err.Error(),
		}
		err = nil
		return result, err
	}

	var posts []*entity.Post

	foundPostErr := p.PostRepository.Get(transaction.Tx, &posts, request.Order, request.Limit, request.Offset)
	if foundPostErr != nil {
		transaction.TxErr = foundPostErr
		result = nil
		err = foundPostErr
		return result, err
	}

	result = &response.Response[[]*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    converter.PostToResponses(posts),
	}
	err = nil
	return result, err
}

func (p PostUseCase) Update(ctx context.Context, request *model_controller.UpdatePostRequest) (result *response.Response[*response.PostResponse], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	countedPost, countedPostErr := p.PostRepository.CountByID(transaction.Tx, request.ID)

	if countedPostErr != nil {
		transaction.TxErr = countedPostErr
		result = nil
		err = countedPostErr
		return result, err
	}

	if countedPost == 0 {
		result = &response.Response[*response.PostResponse]{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
			Data:    nil,
		}
		err = nil
		return result, err
	}

	post := &entity.Post{
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
		UpdatedAt:   null.NewTime(time.Now(), true),
	}

	updatedPostErr := p.PostRepository.Update(transaction.Tx, post, request.ID)
	if updatedPostErr != nil {
		transaction.TxErr = updatedPostErr
		result = nil
		err = updatedPostErr
		return result, err
	}

	result = &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}
	err = nil
	return result, err
}

func (p PostUseCase) Delete(ctx context.Context, request *model_controller.DeletePostRequest) (result *response.Response[*response.PostResponse], err error) {
	transaction := ctx.Value("transaction").(*model.Transaction)

	countedPost, countedPostErr := p.PostRepository.CountByID(transaction.Tx, request.ID)

	if countedPostErr != nil {
		transaction.TxErr = countedPostErr
		result = nil
		err = countedPostErr
		return result, err
	}

	if countedPost == 0 {
		result = &response.Response[*response.PostResponse]{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
			Data:    nil,
		}
		err = nil
		return result, err
	}

	_, deletedPostErr := p.PostRepository.Delete(transaction.Tx, request.ID)
	if deletedPostErr != nil {
		transaction.TxErr = deletedPostErr
		result = nil
		err = deletedPostErr
		return result, err
	}

	result = &response.Response[*response.PostResponse]{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	}
	err = nil
	return result, err
}
