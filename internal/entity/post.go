package entity

import "github.com/guregu/null"

type Post struct {
	Id          null.String `json:"id" bson:"id"`
	User_Id     null.String `json:"user_id" bson:"user_id"`
	Description null.String `json:"description" bson:"description"`
	Image       null.String `json:"image" bson:"image"`
	UpdatedAt   null.Time   `json:"updated_at" bson:"updated_at"`
	CreatedAt   null.Time   `json:"created_at" bson:"created_at"`
	DeletedAt   null.Time   `json:"deleted_at" bson:"deleted_at"`
}
