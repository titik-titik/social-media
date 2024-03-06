package seeder

import (
	"social-media/internal/repository"
	"social-media/test/mock"
)

type User struct {
	UserMock       *mock.UserMock
	UserRepository *repository.UserRepository
}

func NewUser(userRepository *repository.UserRepository) *User {
	userSeeder := &User{
		UserMock:       mock.NewUserMock(),
		UserRepository: userRepository,
	}
	return userSeeder
}

func (userSeeder *User) Up() {
	for _, user := range userSeeder.UserMock.Data {
		userSeeder.UserRepository.CreateOne(user)
	}
}

func (userSeeder *User) Down() {
	for _, user := range userSeeder.UserMock.Data {
		userSeeder.UserRepository.DeleteOneById(user.Id.String)
	}
}
