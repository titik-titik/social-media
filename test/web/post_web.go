package web

import (
	"log"
	"testing"
)

type PostWeb struct {
	Test *testing.T
	Path string
}

func NewPostWeb(test *testing.T) *PostWeb {
	postWeb := &PostWeb{
		Test: test,
		Path: "posts",
	}
	return postWeb
}

func (p *PostWeb) Start() {
	p.Test.Run("GetAllPost", p.GetAllPost)
}

func (p PostWeb) GetAllPost(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	log.Fatalf("post %+v", testWeb.AllSeeder.Post)
}
