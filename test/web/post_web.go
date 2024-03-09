package web

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	model_response "social-media/internal/model/response"
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
	p.Test.Run("FindByID", p.FindByID)
}

func (p PostWeb) FindByID(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	postMock := testWeb.AllSeeder.Post.PostMock.Data[0]
	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, p.Path, postMock.Id.String)

	request, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Fatal(err)
	}

	bodyResponse := &model_response.Response[*model_response.PostResponse]{}
	if err := json.NewDecoder(response.Body).Decode(bodyResponse); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, postMock.Id, bodyResponse.Data.ID)
	assert.Equal(t, postMock.UserId, bodyResponse.Data.UserId)
	assert.Equal(t, postMock.ImageUrl, bodyResponse.Data.ImageUrl)
	assert.Equal(t, postMock.Description, bodyResponse.Data.Description)
}
