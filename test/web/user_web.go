package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/guregu/null"
	"github.com/stretchr/testify/assert"
	"net/http"
	"social-media/internal/entity"
	"social-media/internal/model"
	"social-media/internal/model/request/user"
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
	userWeb.Test.Run("UserWeb PatchOneById", userWeb.PatchOneById)
	userWeb.Test.Run("UserWeb DeleteOneById", userWeb.DeleteOneById)
}

func (userWeb *UserWeb) FindOneById(t *testing.T) {
	// t.Parallel()

	allSeeder := testWeb.GetAllSeeder()
	allSeeder.Up()
	t.Cleanup(allSeeder.Down)

	selectedUserData := allSeeder.UserSeeder.UserMock.Data[0]
	idValue, idValueErr := selectedUserData.Id.Value()
	if idValueErr != nil {
		t.Error(idValueErr)
	}

	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, userWeb.Path, idValue)
	request, newRequestErr := http.NewRequest(http.MethodGet, url, nil)
	if newRequestErr != nil {
		t.Error(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if newRequestErr != nil {
		t.Error(newRequestErr)
	}
	if doErr != nil {
		t.Error(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Error(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}

func (userWeb *UserWeb) FindOneByEmail(t *testing.T) {
	// t.Parallel()

	allSeeder := testWeb.GetAllSeeder()
	allSeeder.Up()
	t.Cleanup(allSeeder.Down)

	selectedUserData := allSeeder.UserSeeder.UserMock.Data[0]
	email, valueErr := selectedUserData.Email.Value()
	if valueErr != nil {
		t.Error(valueErr)
	}

	url := fmt.Sprintf("%s/%s?email=%s", testWeb.Server.URL, userWeb.Path, email)
	request, newRequestErr := http.NewRequest(http.MethodGet, url, nil)
	if newRequestErr != nil {
		t.Error(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Error(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Error(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}

func (userWeb *UserWeb) FindOneByUsername(t *testing.T) {
	// t.Parallel()

	allSeeder := testWeb.GetAllSeeder()
	allSeeder.Up()
	t.Cleanup(allSeeder.Down)

	selectedUserData := allSeeder.UserSeeder.UserMock.Data[0]
	username, valueErr := selectedUserData.Username.Value()
	if valueErr != nil {
		t.Error(valueErr)
	}

	url := fmt.Sprintf("%s/%s?username=%s", testWeb.Server.URL, userWeb.Path, username)
	request, newRequestErr := http.NewRequest(http.MethodGet, url, nil)
	if newRequestErr != nil {
		t.Error(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Error(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Error(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}

func (userWeb *UserWeb) PatchOneById(t *testing.T) {
	// t.Parallel()

	allSeeder := testWeb.GetAllSeeder()
	allSeeder.Up()
	t.Cleanup(allSeeder.Down)

	selectedUserData := allSeeder.UserSeeder.UserMock.Data[0]
	idValue, idValueErr := selectedUserData.Id.Value()
	if idValueErr != nil {
		t.Error(idValueErr)
	}

	bodyRequest := &user.PatchOneById{}
	nameValue, nameValueErr := selectedUserData.Name.Value()
	if nameValueErr != nil {
		t.Error(nameValueErr)
	}
	bodyRequest.Name = null.NewString(nameValue.(string)+"patched", true)
	emailValue, emailValueErr := selectedUserData.Email.Value()
	if emailValueErr != nil {
		t.Error(emailValueErr)
	}
	bodyRequest.Email = null.NewString(emailValue.(string)+"patched", true)
	usernameValue, usernameValueErr := selectedUserData.Username.Value()
	if usernameValueErr != nil {
		t.Error(usernameValueErr)
	}
	bodyRequest.Username = null.NewString(usernameValue.(string)+"patched", true)
	passwordValue, passwordValueErr := selectedUserData.Password.Value()
	if passwordValueErr != nil {
		t.Error(passwordValueErr)
	}
	bodyRequest.Password = null.NewString(passwordValue.(string)+"patched", true)
	avatarUrlValue, avatarUrlValueErr := selectedUserData.AvatarUrl.Value()
	if avatarUrlValueErr != nil {
		t.Error(avatarUrlValueErr)
	}
	bodyRequest.AvatarUrl = null.NewString(avatarUrlValue.(string)+"patched", true)
	bioValue, bioValueErr := selectedUserData.Bio.Value()
	if bioValueErr != nil {
		t.Error(bioValueErr)
	}
	bodyRequest.Bio = null.NewString(bioValue.(string)+"patched", true)

	bodyRequestJsonByte, marshalErr := json.Marshal(bodyRequest)
	if marshalErr != nil {
		t.Error(marshalErr)
	}
	bodyRequestBuffer := bytes.NewBuffer(bodyRequestJsonByte)

	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, userWeb.Path, idValue)
	request, newRequestErr := http.NewRequest(http.MethodPatch, url, bodyRequestBuffer)
	if newRequestErr != nil {
		t.Error(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Error(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Error(decodeErr)
	}

	assert.Equal(t, bodyRequest.Name, bodyResponse.Data.Name)
	assert.Equal(t, bodyRequest.Email, bodyResponse.Data.Email)
	assert.Equal(t, bodyRequest.Username, bodyResponse.Data.Username)
	assert.Equal(t, bodyRequest.Password, bodyResponse.Data.Password)
	assert.Equal(t, bodyRequest.AvatarUrl, bodyResponse.Data.AvatarUrl)
	assert.Equal(t, bodyRequest.Bio, bodyResponse.Data.Bio)
}

func (userWeb *UserWeb) DeleteOneById(t *testing.T) {
	// t.Parallel()

	allSeeder := testWeb.GetAllSeeder()
	allSeeder.Up()
	t.Cleanup(allSeeder.Down)

	selectedUserData := allSeeder.UserSeeder.UserMock.Data[0]
	idValue, idValueErr := selectedUserData.Id.Value()
	if idValueErr != nil {
		t.Error(idValueErr)
	}

	url := fmt.Sprintf("%s/%s/%s", testWeb.Server.URL, userWeb.Path, idValue)
	request, newRequestErr := http.NewRequest(http.MethodDelete, url, nil)
	if newRequestErr != nil {
		t.Error(newRequestErr)
	}
	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Error(doErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))

	bodyResponse := &model.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Error(decodeErr)
	}

	assert.Equal(t, selectedUserData, bodyResponse.Data)
}
