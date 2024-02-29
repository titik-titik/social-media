package main

import (
	"fmt"
	"net/http"
	"social-media/internal/config"
	http_delivery "social-media/internal/delivery/http"
	"social-media/internal/delivery/http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Web started.")

	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)

	userRepository := repository.NewUserRepository(databaseConfig)
	searchRepository := repository.NewSearchRepository(databaseConfig)
	repositoryConfig := config.NewRepositoryConfig(
		userRepository,
		searchRepository,
	)

	userUseCase := use_case.NewUserUseCase(repositoryConfig)
	searchUseCase := use_case.NewSearchUseCase(repositoryConfig)
	useCaseConfig := config.NewUseCaseConfig(
		userUseCase,
		searchUseCase,
	)

	userController := http_delivery.NewUserController(useCaseConfig)
	searchController := http_delivery.NewSearchController(useCaseConfig)
	controllerConfig := config.NewControllerConfig(
		userController,
		searchController,
	)

	router := mux.NewRouter()
	userRoute := route.NewUserRoute(router, controllerConfig)
	rootRoute := route.NewRootRoute(
		router,
		userRoute,
	)

	rootRoute.Register()

	address := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.Port,
	)
	listenAndServeErr := http.ListenAndServe(address, rootRoute.Router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}

	fmt.Println("Web finished.")
}
