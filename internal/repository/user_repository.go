package repository

import "social-media/internal/entity"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	userRepository := &UserRepository{}
	return userRepository
}

func (userRepository *UserRepository) FindOneById(id string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByUsername(username string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByEmail(email string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByEmailAndPassword(email string, password string) *entity.User {
	return nil
}

func (userRepository *UserRepository) FindOneByUsernameAndPassword(username string, password string) *entity.User {
	return nil
}

func (userRepository *UserRepository) CreateOne(toCreateUser *entity.User) *entity.User {
	return nil
}

func (userRepository *UserRepository) PatchOneById(id string, toPatchUser *entity.User) *entity.User {
	return nil
}

func (userRepository *UserRepository) DeleteOneById(id string) *entity.User {
	return nil
}
