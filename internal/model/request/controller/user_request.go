package request

import "github.com/guregu/null"

type UserPatchOneByIdRequest struct {
	Name      null.String `json:"name"`
	Username  null.String `json:"username"`
	Email     null.String `json:"email"`
	Password  null.String `json:"password"`
	AvatarUrl null.String `json:"avatar_url"`
	Bio       null.String `json:"bio"`
}
