package entity

import (
	"github.com/guregu/null"
)

type Session struct {
	ID                    null.String `json:"id" bson:"id"`
	UserID                null.String `json:"user_id" bson:"user_id"`
	AccessToken           null.String `json:"access_token" bson:"access_token"`
	RefreshToken          null.String `json:"refresh_token" bson:"refresh_token"`
	AccessTokenExpiredAt  null.Time   `json:"access_token_expired_at" bson:"access_token_expired_at"`
	RefreshTokenExpiredAt null.Time   `json:"refresh_token_expired_at" bson:"refresh_token_expired_at"`
	UpdatedAt             null.Time   `json:"updated_at" bson:"updated_at"`
	CreatedAt             null.Time   `json:"created_at" bson:"created_at"`
	DeletedAt             null.Time   `json:"deleted_at" bson:"deleted_at"`
}
