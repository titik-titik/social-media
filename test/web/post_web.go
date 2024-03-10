package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
	"net/http"
	model_request "social-media/internal/model/request/controller"
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
	p.Test.Run("GetAll", p.GetAll)
	p.Test.Run("UpdatePost", p.Update)
	p.Test.Run("DeletePost", p.Delete)
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

func (p PostWeb) GetAll(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	url := fmt.Sprintf("%s/%s/", testWeb.Server.URL, p.Path)

	jsonBody := []byte(`{
    "limit": 10,
    "offset": 0,
    "order":"DESC"
}`)
	bodyReader := bytes.NewReader(jsonBody)
	request, err := http.NewRequest(http.MethodGet, url, bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Fatal(err)
	}

	bodyResponse := model_response.Response[[]*model_response.PostResponse]{}
	if err := json.NewDecoder(response.Body).Decode(&bodyResponse); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, bodyResponse.Code, http.StatusOK)
	assert.Equal(t, bodyResponse.Message, http.StatusText(http.StatusOK))
	assert.Len(t, bodyResponse.Data, 10)
}

func (p PostWeb) Update(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	postMock := testWeb.AllSeeder.Post.PostMock.Data[0]
	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, p.Path, postMock.Id.String)

	newPost := model_request.UpdatePostRequest{
		ID:          postMock.Id.String,
		ImageUrl:    null.NewString("http://test-update.com/image.jpg", true),
		Description: null.NewString("Test Update Description", true),
	}

	newPostByte, err := json.Marshal(newPost)
	if err != nil {
		t.Fatal(err)
	}
	newPostReqBuffer := bytes.NewBuffer(newPostByte)

	request, err := http.NewRequest(http.MethodPut, url, newPostReqBuffer)
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

	assert.Equal(t, bodyResponse.Code, http.StatusOK)
	assert.Equal(t, bodyResponse.Message, http.StatusText(http.StatusOK))

	url = fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, p.Path, postMock.Id.String)

	request, err = http.NewRequest(http.MethodGet, url, http.NoBody)

	if err != nil {
		t.Fatal(err)
	}

	response, err = http.DefaultClient.Do(request)

	if err != nil {
		t.Fatal(err)
	}

	bodyResponseGetUpdate := &model_response.Response[*model_response.PostResponse]{}
	if err := json.NewDecoder(response.Body).Decode(bodyResponseGetUpdate); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, bodyResponse.Code, http.StatusOK)
	assert.Equal(t, bodyResponse.Message, http.StatusText(http.StatusOK))
	assert.Equal(t, postMock.Id.String, bodyResponseGetUpdate.Data.ID.String)
	assert.Equal(t, postMock.UserId, bodyResponseGetUpdate.Data.UserId)
	assert.Equal(t, newPost.ImageUrl, bodyResponseGetUpdate.Data.ImageUrl)
	assert.Equal(t, newPost.Description, bodyResponseGetUpdate.Data.Description)
	assert.True(t, postMock.UpdatedAt.Time.Before(bodyResponseGetUpdate.Data.UpdatedAt.Time))
}

func (p PostWeb) Delete(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	postMock := testWeb.AllSeeder.Post.PostMock.Data[1]
	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, p.Path, postMock.Id.String)

	request, err := http.NewRequest(http.MethodDelete, url, http.NoBody)
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
	assert.Equal(t, bodyResponse.Code, http.StatusOK)
	assert.Equal(t, bodyResponse.Message, http.StatusText(http.StatusOK))
	assert.Empty(t, bodyResponse.Data)
}
