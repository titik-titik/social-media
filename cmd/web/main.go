package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"social-media/internal/config"
	http_delivery "social-media/internal/delivery/http"
	"social-media/internal/delivery/http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"
)

func main() {
	fmt.Println("Web started.")

	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		panic(fmt.Errorf("error loading .env file: %w", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)

	searchRepository := repository.NewSearchRepository(databaseConfig)
	userRepository := repository.NewUserRepository(databaseConfig)

	userUseCase := use_case.NewUserUseCase(userRepository)
	searchUseCase := use_case.NewSearchUseCase(searchRepository)

	userController := http_delivery.NewUserController(userUseCase, searchUseCase)

	router := mux.NewRouter()
	userRoute := route.NewUserRoute(router, userController)
	rootRoute := route.NewRootRoute(
		userRoute,
	)

	rootRoute.Register()

	address := fmt.Sprintf(
		"%s:%s",
		envConfig.App.Host,
		envConfig.App.Port,
	)
	listenAndServeErr := http.ListenAndServe(address, router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}

	fmt.Println("Web finished.")
}
