package container

import (
	"fmt"
	"social-media/internal/config"
	http_delivery "social-media/internal/delivery/http"
	"social-media/internal/delivery/http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type WebContainer struct {
	Env        *config.EnvConfig
	Database   *config.DatabaseConfig
	Repository *RepositoryContainer
	UseCase    *UseCaseContainer
	Controller *ControllerContainer
	Route      *route.RootRoute
}

func NewWebContainer() *WebContainer {
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)

	searchRepository := repository.NewSearchRepository()
	userRepository := repository.NewUserRepository()
	sessionRepository := repository.NewSessionRepository()
	postRepository := repository.NewPostRepository()
	repositoryContainer := NewRepositoryContainer(userRepository, sessionRepository, postRepository, searchRepository)

	searchUseCase := use_case.NewSearchUseCase(databaseConfig, searchRepository)
	userUseCase := use_case.NewUserUseCase(databaseConfig, userRepository)
	authUseCase := use_case.NewAuthUseCase(databaseConfig, userRepository, sessionRepository)
	postUseCase := use_case.NewPostUseCase(databaseConfig, postRepository)
	useCaseContainer := NewUseCaseContainer(userUseCase, authUseCase, searchUseCase)

	userController := http_delivery.NewUserController(userUseCase)
	postController := http_delivery.NewPostController(postUseCase)
	authController := http_delivery.NewAuthController(authUseCase)
	controllerContainer := NewControllerContainer(userController, authController)

	router := mux.NewRouter()

	userRoute := route.NewUserRoute(router, userController)
	authRoute := route.NewAuthRoute(router, authController)
	postRoute := route.NewPostRoute(router, postController)

	rootRoute := route.NewRootRoute(
		router,
		userRoute,
		authRoute,
		postRoute,
	)

	rootRoute.Register()

	webContainer := &WebContainer{
		Env:        envConfig,
		Database:   databaseConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Controller: controllerContainer,
		Route:      rootRoute,
	}

	return webContainer
}
