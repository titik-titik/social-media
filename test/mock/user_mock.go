package mock

import (
	"social-media/internal/entity"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

type UserMock struct {
	Data []*entity.User
}

func NewUserMock() *UserMock {
	currentTime := time.Now().UTC()

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
				CreatedAt:  null.NewTime(currentTime.Add(0*time.Second), true),
				UpdatedAt:  null.NewTime(currentTime.Add(0*time.Second), true),
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
				CreatedAt:  null.NewTime(currentTime.Add(1*time.Second), true),
				UpdatedAt:  null.NewTime(currentTime.Add(1*time.Second), true),
				DeletedAt:  null.NewTime(time.Time{}, false),
			},
		},
	}
	return userMock
}
