package request

import "github.com/guregu/null"

type CreatePostRequest struct {
	ImageUrl    null.String `json:"image_url"`
	Description null.String `json:"description"`
}

type GetPostRequest struct {
	PostId null.String `json:"post_id"`
}

type GetAllPostRequest struct {
	Limit  int8   `json:"limit" validate:"required"`
	Offset int64  `json:"offset" validate:"required"`
	Order  string `json:"order" validate:"required"`
}

type UpdatePostRequest struct {
	ID          string      `json:"-"`
	ImageUrl    null.String `json:"image_url"`
	Description null.String `json:"description"`
}

type DeletePostRequest struct {
	ID string `json:"-"`
}
