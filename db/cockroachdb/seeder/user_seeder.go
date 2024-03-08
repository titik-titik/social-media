package seeder

import (
	"social-media/internal/repository"
	"social-media/test/mock"
	"social-media/tool"

	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
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
		userCopied := tool.DeepCopy(user)
		hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(userCopied.Password.String), bcrypt.DefaultCost)
		if hashedPasswordErr != nil {
			panic(hashedPasswordErr)
		}
		userCopied.Password = null.NewString(string(hashedPassword), true)
		userSeeder.UserRepository.CreateOne(userCopied)
	}
}

func (userSeeder *UserSeeder) Down() {
	for _, user := range userSeeder.UserMock.Data {
		userSeeder.UserRepository.DeleteOneById(user.Id.String)
	}
}
