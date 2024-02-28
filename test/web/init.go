package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http/httptest"
	"social-media/internal/config"
	http_delivery "social-media/internal/delivery/http"
	"social-media/internal/delivery/http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"
	"social-media/seeder"
)

var testWeb *TestWeb

type TestWeb struct {
	Server         *httptest.Server
	UserWeb        *UserWeb
	UserRepository *repository.UserRepository
}

func NewTestWeb() *TestWeb {
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
		userRoute,
	)
	rootRoute.Register()
	server := httptest.NewServer(router)

	testWeb := &TestWeb{
		Server:         server,
		UserRepository: userRepository,
	}

	return testWeb
}

func (web *TestWeb) GetAllSeeder() *seeder.AllSeeder {
	userSeeder := seeder.NewUserSeeder(web.UserRepository)
	allSeeder := seeder.NewAllSeeder(userSeeder)
	return allSeeder
}

func init() {
	testWeb = NewTestWeb()
}
