package use_case

import (
	"social-media/internal/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	AuthRepository repository.AuthRepository
}

func NewAuthUseCase(AuthRepository repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		AuthRepository: AuthRepository,
	}
}

func (c *AuthUseCase) Register(username, password, email, avatarURL, bio string) error {
	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate a UUID if ID is not provided
	uuid := uuid.New()
	userID := uuid.String()

	// Current time
	currentTime := time.Now()

	// Save the user to the database using the validated data
	err = c.AuthRepository.Register(userID, username, string(hashedPassword), email, avatarURL, bio, false, currentTime, currentTime)
	if err != nil {
		return err
	}

	return nil
}
