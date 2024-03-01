package web

import (
	"fmt"
	"net/http/httptest"
	seeder "social-media/db/cockroachdb_one/seeder"
	"social-media/internal/config"
	http_delivery "social-media/internal/delivery/delivery_http"
	"social-media/internal/delivery/delivery_http/route"
	"social-media/internal/repository"
	"social-media/internal/use_case"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type TestWeb struct {
	Server         *httptest.Server
	UserRepository *repository.UserRepository
	AllSeeder      *seeder.AllSeeder
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
		router,
		userRoute,
	)
	rootRoute.Register()

	server := httptest.NewServer(rootRoute.Router)

	testWeb := &TestWeb{
		Server:         server,
		UserRepository: userRepository,
	}

	return testWeb
}

func (web *TestWeb) GetAllSeeder() *seeder.AllSeeder {
	userSeeder := seeder.NewUserSeeder(web.UserRepository)
	seederConfig := seeder.NewAllSeeder(
		userSeeder,
	)
	return seederConfig
}

func GetTestWeb() *TestWeb {
	testWeb := NewTestWeb()
	testWeb.AllSeeder = testWeb.GetAllSeeder()
	return testWeb
}
