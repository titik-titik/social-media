package response

import "github.com/guregu/null"

type UserResponse struct {
	ID          null.String `json:"id,omitempty"`
	UserId      null.String `json:"user_id,omitempty"`
	ImageUrl    null.String `json:"image_url,omitempty"`
	Description null.String `json:"description,omitempty"`
	CreatedAt   null.Time   `json:"created_at,omitempty"`
	UpdatedAt   null.Time   `json:"updated_at,omitempty"`
}
