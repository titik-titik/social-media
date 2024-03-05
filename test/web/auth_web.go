package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"social-media/internal/entity"
	model_request "social-media/internal/model/request"
	model_response "social-media/internal/model/response"

	"testing"

	"github.com/stretchr/testify/assert"
)

type AuthWeb struct {
	Test *testing.T
	Path string
}

func NewAuthWeb(test *testing.T) *AuthWeb {
	authWeb := &AuthWeb{
		Test: test,
		Path: "auths",
	}
	return authWeb
}
func (authWeb *AuthWeb) Start() {
	authWeb.Test.Run("authWeb Register", authWeb.Register)
}
func (authWeb *AuthWeb) Register(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	mockUser := testWeb.AllSeeder.UserSeeder.UserMock.Data[0]

	bodyRequest := &model_request.RegisterRequest{
		Username: mockUser.Username,
		Email:    mockUser.Email,
		Password: mockUser.Password,
	}

	bodyRequestJsonByte, marshalErr := json.Marshal(bodyRequest)
	if marshalErr != nil {
		t.Fatal(marshalErr)
	}
	bodyRequestBuffer := bytes.NewBuffer(bodyRequestJsonByte)

	url := fmt.Sprintf("%s/%s", testWeb.Server.URL, authWeb.Path)
	request, newRequestErr := http.NewRequest(http.MethodPost, url, bodyRequestBuffer)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}

	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, mockUser.Username, bodyResponse.Data.Username.String)
	assert.Equal(t, mockUser.Email, bodyResponse.Data.Email.String)
}
