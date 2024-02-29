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

	userRepository := repository.NewUserRepository(databaseConfig)
	userUseCase := use_case.NewUserUseCase(userRepository)

	userController := http_delivery.NewUserController(userUseCase)

	router := mux.NewRouter()
	userRoute := route.NewUserRoute(router, userController)
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
