package container

import (
	"fmt"
	"social-media/internal/config"
	httpdelivery "social-media/internal/delivery/delivery_http"
	"social-media/internal/delivery/delivery_http/route"
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

	// Setup Config
	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)

	// Setup Repository
	searchRepository := repository.NewSearchRepository(databaseConfig)
	userRepository := repository.NewUserRepository(databaseConfig)
	// postRepository := repository.NewPostRepository()
	repositoryContainer := NewRepositoryContainer(userRepository, searchRepository)
	// Setup UseCase
	searchUseCase := use_case.NewSearchUseCase(searchRepository)
	userUseCase := use_case.NewUserUseCase(userRepository)
	// postUseCase := use_case.NewPostUseCase(databaseConfig, postRepository)
	useCaseContainer := NewUseCaseContainer(userUseCase, searchUseCase)
	// Setup Controller
	userController := httpdelivery.NewUserController(userUseCase)
	// postController := httpdelivery.NewPostController(postUseCase)
	controllerContainer := NewControllerContainer(userController)

	// init router
	router := mux.NewRouter()

	//define routes
	userRoute := route.NewUserRoute(router, userController)
	// postRoute := route.NewPostRoute(router, postController)

	//bootstrap route
	rootRoute := route.NewRootRoute(
		router,
		userRoute,
		// postRoute,
	)

	rootRoute.Register()

	return &WebContainer{
		Env:        envConfig,
		Database:   databaseConfig,
		Repository: repositoryContainer,
		UseCase:    useCaseContainer,
		Controller: controllerContainer,
		Route:      rootRoute,
	}
}
