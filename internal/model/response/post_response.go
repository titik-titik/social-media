package response

import "github.com/guregu/null"

type PostResponse struct {
	UserId      null.String `json:"user_id"`
	ImageUrl    null.String `json:"image_url"`
	Description null.String `json:"description"`
	CreatedAt   null.Time   `json:"created_at"`
	UpdatedAt   null.Time   `json:"updated_at"`
}
