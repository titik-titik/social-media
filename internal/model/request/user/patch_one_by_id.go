package user

import "github.com/guregu/null"

type PatchOneById struct {
	Name      null.String `json:"name"`
	Username  null.String `json:"username"`
	Email     null.String `json:"email"`
	Password  null.String `json:"password"`
	AvatarUrl null.String `json:"avatar_url"`
	Bio       null.String `json:"bio"`
}
