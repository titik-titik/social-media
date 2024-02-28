package entity

import (
	"github.com/guregu/null"
)

type User struct {
	Id         null.String `json:"id" bson:"id"`
	Name       null.String `json:"name" bson:"name"`
	Username   null.String `json:"username" bson:"username"`
	Email      null.String `json:"email" bson:"email"`
	Password   null.String `json:"password" bson:"password"`
	AvatarUrl  null.String `json:"avatar_url" bson:"avatar_url"`
	Bio        null.String `json:"bio" bson:"bio"`
	IsVerified null.Bool   `json:"is_verified" bson:"is_verified"`
	UpdatedAt  null.Time   `json:"updated_at" bson:"updated_at"`
	CreatedAt  null.Time   `json:"created_at" bson:"created_at"`
	DeletedAt  null.Time   `json:"deleted_at" bson:"deleted_at"`
}
