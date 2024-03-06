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

	// Load Environment Variable
	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)

	searchRepository := repository.NewSearchRepository(databaseConfig)
	userRepository := repository.NewUserRepository(databaseConfig)
	postRepository := repository.NewPostRepository()
	authRepository := repository.NewAuthRepository(databaseConfig)
	repositoryContainer := NewRepositoryContainer(userRepository, postRepository, authRepository, searchRepository)

	searchUseCase := use_case.NewSearchUseCase(searchRepository)
	userUseCase := use_case.NewUserUseCase(userRepository)
	authUseCase := use_case.NewAuthUseCase(authRepository)
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
