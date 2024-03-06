package web

import (
	"net/http/httptest"
	"social-media/container"
	seeder "social-media/db/cockroachdb_one/seeder"
)

type TestWeb struct {
	Server    *httptest.Server
	AllSeeder *seeder.AllSeeder
	Container *container.WebContainer
}

func NewTestWeb() *TestWeb {
	webContainer := container.NewWebContainer()

	server := httptest.NewServer(webContainer.Route.Router)

	testWeb := &TestWeb{
		Server:    server,
		Container: webContainer,
	}

	return testWeb
}

func (web *TestWeb) GetAllSeeder() *seeder.AllSeeder {
	userSeeder := seeder.NewUserSeeder(web.Container.Repository.User)
	AuthSeeder := seeder.NewAuthSeeder(web.Container.Repository.Auth)
	seederConfig := seeder.NewAllSeeder(
		userSeeder,
		AuthSeeder,
	)
	return seederConfig
}

func GetTestWeb() *TestWeb {
	testWeb := NewTestWeb()
	testWeb.AllSeeder = testWeb.GetAllSeeder()
	return testWeb
}
