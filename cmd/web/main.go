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
	fmt.Println("Web starting.")

	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		fmt.Println(fmt.Sprintf("Error loading .env file: %+v", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	databaseConfig := config.NewDatabaseConfig(envConfig)
	userRepository := repository.NewUserRepository(databaseConfig)
	userUseCase := use_case.NewUserUseCase(userRepository)
	userController := http_delivery.NewUserController(userUseCase)

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
