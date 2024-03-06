package seeder

import (
	"social-media/internal/repository"
	"social-media/test/mock"
)

type AuthSeeder struct {
	AuthMock       *mock.AuthMock
	AuthRepository *repository.AuthRepository
}

func NewAuthSeeder(AuthRepository *repository.AuthRepository) *AuthSeeder {
	AuthSeeder := &AuthSeeder{
		AuthMock:       mock.NewAuthMock(),
		AuthRepository: AuthRepository,
	}
	return AuthSeeder
}

func (AuthSeeder *AuthSeeder) Up() {
	for _, Auth := range AuthSeeder.AuthMock.Data {
		AuthSeeder.AuthRepository.CreateDummy(Auth)
	}
}

func (AuthSeeder *AuthSeeder) Down() {
	for _, Auth := range AuthSeeder.AuthMock.Data {
		AuthSeeder.AuthRepository.DeleteDummyById(Auth.Id.String)
	}
}
