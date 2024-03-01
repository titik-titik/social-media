package main

import (
	"fmt"
	"net/http"
	"social-media/internal/config"
	httpdelivery "social-media/internal/delivery/delivery_http"
	"social-media/internal/delivery/delivery_http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Web started.")

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
	postRepository := repository.NewPostRepository()

	// Setup UseCase
	_ = use_case.NewSearchUseCase(searchRepository)
	userUseCase := use_case.NewUserUseCase(userRepository)
	postUseCase := use_case.NewPostUseCase(databaseConfig, postRepository)

	// Setup Controller
	userController := httpdelivery.NewUserController(userUseCase)
	postController := httpdelivery.NewPostController(postUseCase)

	// init router
	router := mux.NewRouter()

	//define routes
	userRoute := route.NewUserRoute(router, userController)
	postRoute := route.NewPostRoute(router, postController)

	//bootstrap route
	rootRoute := route.NewRootRoute(
		router,
		userRoute,
		postRoute,
	)

	rootRoute.Register()

	// Setup Engine
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
