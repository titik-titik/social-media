package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"social-media/internal/entity"
	model_request "social-media/internal/model/request"
	model_response "social-media/internal/model/response"
	"testing"
)

type UserWeb struct {
	Test *testing.T
	Path string
}

func NewUserWeb(test *testing.T) *UserWeb {
	userWeb := &UserWeb{
		Test: test,
		Path: "users",
	}
	return userWeb
}

func (userWeb *UserWeb) Start() {
	userWeb.Test.Run("UserWeb FindOneById", userWeb.FindOneById)
	userWeb.Test.Run("UserWeb FindOneByEmail", userWeb.FindOneByEmail)
	userWeb.Test.Run("UserWeb FindOneByUsername", userWeb.FindOneByUsername)
	userWeb.Test.Run("UserWeb UserPatchOneByIdRequest", userWeb.PatchOneById)
	userWeb.Test.Run("UserWeb DeleteOneById", userWeb.DeleteOneById)
}

func (userWeb *UserWeb) FindOneById(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedUserData := testWeb.AllSeeder.UserSeeder.UserMock.Data[0]

	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, userWeb.Path, selectedUserData.Id.String)
	request, newRequestErr := http.NewRequest(http.MethodGet, url, http.NoBody)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}
	if doErr != nil {
		t.Fatal(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}

func (userWeb *UserWeb) FindOneByEmail(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedUserData := testWeb.AllSeeder.UserSeeder.UserMock.Data[0]

	url := fmt.Sprintf("%s/%s?email=%s", testWeb.Server.URL, userWeb.Path, selectedUserData.Email.String)
	request, newRequestErr := http.NewRequest(http.MethodGet, url, http.NoBody)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}

func (userWeb *UserWeb) FindOneByUsername(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedUserData := testWeb.AllSeeder.UserSeeder.UserMock.Data[0]

	url := fmt.Sprintf("%s/%s?username=%s", testWeb.Server.URL, userWeb.Path, selectedUserData.Username.String)
	request, newRequestErr := http.NewRequest(http.MethodGet, url, http.NoBody)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}

func (userWeb *UserWeb) PatchOneById(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedUserData := testWeb.AllSeeder.UserSeeder.UserMock.Data[0]

	bodyRequest := &model_request.UserPatchOneByIdRequest{}
	bodyRequest.Name = null.NewString(selectedUserData.Name.String+"patched", true)
	bodyRequest.Email = null.NewString(selectedUserData.Email.String+"patched", true)
	bodyRequest.Username = null.NewString(selectedUserData.Username.String+"patched", true)
	bodyRequest.Password = null.NewString(selectedUserData.Password.String+"patched", true)
	bodyRequest.AvatarUrl = null.NewString(selectedUserData.AvatarUrl.String+"patched", true)
	bodyRequest.Bio = null.NewString(selectedUserData.Bio.String+"patched", true)

	bodyRequestJsonByte, marshalErr := json.Marshal(bodyRequest)
	if marshalErr != nil {
		t.Fatal(marshalErr)
	}
	bodyRequestBuffer := bytes.NewBuffer(bodyRequestJsonByte)

	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, userWeb.Path, selectedUserData.Id.String)
	request, newRequestErr := http.NewRequest(http.MethodPatch, url, bodyRequestBuffer)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, selectedUserData.Id, bodyResponse.Data.Id)
	assert.Equal(t, bodyRequest.Name, bodyResponse.Data.Name)
	assert.Equal(t, bodyRequest.Email, bodyResponse.Data.Email)
	assert.Equal(t, bodyRequest.Username, bodyResponse.Data.Username)
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(bodyResponse.Data.Password.String), []byte(bodyRequest.Password.String)))
	assert.Equal(t, bodyRequest.AvatarUrl, bodyResponse.Data.AvatarUrl)
	assert.Equal(t, bodyRequest.Bio, bodyResponse.Data.Bio)
}

func (userWeb *UserWeb) DeleteOneById(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedUserData := testWeb.AllSeeder.UserSeeder.UserMock.Data[0]

	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, userWeb.Path, selectedUserData.Id.String)
	request, newRequestErr := http.NewRequest(http.MethodDelete, url, http.NoBody)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}
