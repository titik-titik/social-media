package config

import (
	"social-media/internal/delivery/http"
	"social-media/internal/repository"
	"social-media/internal/use_case"
)

type BootstrapConfig struct {
	Env *EnvConfig
}

func Bootstrap(bootstrapConfig *BootstrapConfig) {
	userRepository := repository.NewUserRepository()
	userUseCase := use_case.NewUserUseCase(userRepository)
	_ = http.NewUserController(userUseCase)
}
