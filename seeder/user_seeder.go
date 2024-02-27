package seeder

import (
	"social-media/internal/repository"
	"social-media/test/mock"
)

type UserSeeder struct {
	UserMock       *mock.UserMock
	UserRepository *repository.UserRepository
}

func NewUserSeeder(userRepository *repository.UserRepository) *UserSeeder {
	userSeeder := &UserSeeder{
		UserMock:       mock.NewUserMock(),
		UserRepository: userRepository,
	}
	return userSeeder
}

func (userSeeder *UserSeeder) Up() {
	for _, user := range userSeeder.UserMock.Data {
		userSeeder.UserRepository.CreateOne(user)
	}
}

func (userSeeder *UserSeeder) Down() {
	for _, user := range userSeeder.UserMock.Data {
		id, valueErr := user.Id.Value()
		if valueErr != nil {
			panic(valueErr)
		}
		userSeeder.UserRepository.DeleteOneById(id.(string))
	}
}
