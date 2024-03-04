package request

import "github.com/guregu/null"

type RegisterRequest struct {
	Username null.String `json:"username"`
	Email    null.String `json:"email"`
	Password null.String `json:"password"`
}
