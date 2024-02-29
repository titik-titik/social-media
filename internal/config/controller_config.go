package config

import "social-media/internal/delivery/http"

type ControllerConfig struct {
	userController   *http.UserController
	searchController *http.SearchController
}

func NewControllerConfig(userController *http.UserController, searchController *http.SearchController) *ControllerConfig {
	return &ControllerConfig{userController: userController, searchController: searchController}
}
