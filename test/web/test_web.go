package web

import (
	"net/http/httptest"
	"social-media/container"
	seeder "social-media/db/cockroachdb/seeder"
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
	userSeeder := seeder.NewUserSeeder(web.Container.Database)
	sessionSeeder := seeder.NewSessionSeeder(web.Container.Database, userSeeder)
	postSeeder := seeder.NewPostSeeder(web.Container.Database, userSeeder)
	seederConfig := seeder.NewAllSeeder(
		userSeeder,
		sessionSeeder,
		postSeeder,
	)
	return seederConfig
}

func GetTestWeb() *TestWeb {
	testWeb := NewTestWeb()
	testWeb.AllSeeder = testWeb.GetAllSeeder()
	return testWeb
}
