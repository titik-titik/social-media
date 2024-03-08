package mock

import (
	"social-media/internal/entity"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
)

type SessionMock struct {
	Data     []*entity.Session
	UserMock *UserMock
}

func NewSessionMock(
	userMock *UserMock,
) *SessionMock {
	currentTime := time.Now().UTC()
	sessionMock := &SessionMock{
		UserMock: userMock,
		Data: []*entity.Session{
			{
				Id:                    null.NewString(uuid.NewString(), true),
				UserId:                null.NewString(userMock.Data[0].Id.String, true),
				AccessToken:           null.NewString(uuid.NewString(), true),
				RefreshToken:          null.NewString(uuid.NewString(), true),
				AccessTokenExpiredAt:  null.NewTime(currentTime.Add(time.Minute*10), true),
				RefreshTokenExpiredAt: null.NewTime(currentTime.Add(time.Hour*24*2), true),
				CreatedAt:             null.NewTime(currentTime.Add(time.Second*0), true),
				UpdatedAt:             null.NewTime(currentTime.Add(time.Second*0), true),
				DeletedAt:             null.NewTime(time.Time{}, false),
			},
			{
				Id:                    null.NewString(uuid.NewString(), true),
				UserId:                null.NewString(userMock.Data[1].Id.String, true),
				AccessToken:           null.NewString(uuid.NewString(), true),
				RefreshToken:          null.NewString(uuid.NewString(), true),
				AccessTokenExpiredAt:  null.NewTime(currentTime.Add(time.Minute*10), true),
				RefreshTokenExpiredAt: null.NewTime(currentTime.Add(time.Hour*24*2), true),
				CreatedAt:             null.NewTime(currentTime.Add(time.Second*0), true),
				UpdatedAt:             null.NewTime(currentTime.Add(time.Second*0), true),
				DeletedAt:             null.NewTime(time.Time{}, false),
			},
		},
	}
	return sessionMock
}
