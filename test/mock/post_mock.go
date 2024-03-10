package mock

import (
	"social-media/internal/entity"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

type PostMock struct {
	UserMock *UserMock
	Data     []*entity.Post
}

func NewPostMock(
	userMock *UserMock,
) *PostMock {
	currentTime := time.Now()
	currentTimeRfc3339 := currentTime.Format(time.RFC3339)
	currentTimeFromRfc3339, parseErr := time.Parse(time.RFC3339, currentTimeRfc3339)
	if parseErr != nil {
		panic(parseErr)
	}
	postMock := &PostMock{
		UserMock: userMock,
		Data: []*entity.Post{
			{
				Id:          null.NewString(uuid.NewString(), true),
				UserId:      null.NewString(userMock.Data[0].Id.String, true),
				ImageUrl:    null.NewString("https://placehold.co/400x400?text=image_url0", true),
				Description: null.NewString("description0", true),
				CreatedAt:   null.NewTime(currentTimeFromRfc3339.Add(time.Duration(time.Duration.Seconds(0))), true),
				UpdatedAt:   null.NewTime(currentTimeFromRfc3339.Add(time.Duration(time.Duration.Seconds(0))), true),
				DeletedAt:   null.NewTime(time.Time{}, false),
			},
			{
				Id:          null.NewString(uuid.NewString(), true),
				UserId:      null.NewString(userMock.Data[1].Id.String, true),
				ImageUrl:    null.NewString("https://placehold.co/400x400?text=image_url1", true),
				Description: null.NewString("description1", true),
				CreatedAt:   null.NewTime(currentTimeFromRfc3339.Add(time.Duration(time.Duration.Seconds(1))), true),
				UpdatedAt:   null.NewTime(currentTimeFromRfc3339.Add(time.Duration(time.Duration.Seconds(1))), true),
				DeletedAt:   null.NewTime(time.Time{}, false),
			},
		},
	}
	return postMock
}
