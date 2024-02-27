package entity

import "time"

type User struct {
	Id         string     `json:"id" bson:"id"`
	Name       string     `json:"name" bson:"name"`
	Username   string     `json:"username" bson:"username"`
	Email      string     `json:"email" bson:"email"`
	Password   string     `json:"password" bson:"password"`
	AvatarUrl  string     `json:"avatar_url" bson:"avatar_url"`
	Bio        string     `json:"bio" bson:"bio"`
	IsVerified bool       `json:"is_verified" bson:"is_verified"`
	UpdatedAt  *time.Time `json:"updated_at" bson:"updated_at"`
	CreatedAt  *time.Time `json:"created_at" bson:"created_at"`
	DeletedAt  *time.Time `json:"deleted_at" bson:"deleted_at"`
}
