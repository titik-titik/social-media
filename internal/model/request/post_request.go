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
	UserId null.String `json:"user_id"`
}

type UpdatePostRequest struct {
	ImageUrl    null.String `json:"image_url"`
	Description null.String `json:"description"`
}
