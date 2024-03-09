package request

import "github.com/guregu/null"

type RegisterRequest struct {
	Name     null.String `json:"name"`
	Username null.String `json:"username"`
	Email    null.String `json:"email"`
	Password null.String `json:"password"`
}

type LoginRequest struct {
	Email    null.String `json:"email"`
	Password null.String `json:"password"`
}
