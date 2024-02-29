package config

import "social-media/internal/repository"

type RepositoryConfig struct {
	userRepo   *repository.UserRepository
	searchRepo *repository.SearchRepository
}

func NewRepositoryConfig(userRepo *repository.UserRepository, searchRepo *repository.SearchRepository) *RepositoryConfig {
	return &RepositoryConfig{userRepo: userRepo, searchRepo: searchRepo}
}
