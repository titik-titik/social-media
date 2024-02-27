package mock

import (
	"github.com/google/uuid"
	"github.com/guregu/null"
	"social-media/internal/entity"
	"time"
)

type UserMock struct {
	Data []*entity.User
}

func NewUserMock() *UserMock {
	currentTime := time.Now().UTC()
	currentTimeIso8601String := currentTime.Format(time.RFC3339)
	currentTimeIso8601, parseErr := time.Parse(time.RFC3339, currentTimeIso8601String)
	if parseErr != nil {
		panic(parseErr)
	}

	userMock := &UserMock{
		Data: []*entity.User{
			{
				Id:         null.NewString(uuid.NewString(), true),
				Name:       null.NewString("name0", true),
				Username:   null.NewString("username"+uuid.NewString(), true),
				Email:      null.NewString("email"+uuid.NewString()+"@mail.com", true),
				Password:   null.NewString("password0", true),
				AvatarUrl:  null.NewString("https://placehold.co/400x400?text=avatar_url0", true),
				Bio:        null.NewString("bio0", true),
				IsVerified: null.NewBool(false, true),
				CreatedAt:  null.NewTime(currentTimeIso8601.Add(time.Duration(time.Duration.Seconds(0))), true),
				UpdatedAt:  null.NewTime(currentTimeIso8601.Add(time.Duration(time.Duration.Seconds(0))), true),
				DeletedAt:  null.NewTime(time.Time{}, false),
			},
			{
				Id:         null.NewString(uuid.NewString(), true),
				Name:       null.NewString("name1", true),
				Username:   null.NewString("username"+uuid.NewString(), true),
				Email:      null.NewString("email"+uuid.NewString()+"@mail.com", true),
				Password:   null.NewString("password1", true),
				AvatarUrl:  null.NewString("https://placehold.co/400x400?text=avatar_url1", true),
				Bio:        null.NewString("bio1", true),
				IsVerified: null.NewBool(true, true),
				CreatedAt:  null.NewTime(currentTimeIso8601.Add(time.Duration(time.Duration.Seconds(1))), true),
				UpdatedAt:  null.NewTime(currentTimeIso8601.Add(time.Duration(time.Duration.Seconds(1))), true),
				DeletedAt:  null.NewTime(time.Time{}, false),
			},
		},
	}
	return userMock
}
