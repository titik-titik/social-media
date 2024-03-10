package web

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"social-media/internal/entity"
	model_request "social-media/internal/model/request/controller"
	model_response "social-media/internal/model/response"

	"golang.org/x/crypto/bcrypt"

	"testing"

	"github.com/guregu/null"
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
	authWeb.Test.Run("AuthWeb_Register_Succeed", authWeb.Register)
	authWeb.Test.Run("AuthWeb_Login_Succeed", authWeb.Login)
	authWeb.Test.Run("AuthWeb_Logout_Succeed", authWeb.Logout)
	authWeb.Test.Run("AuthWeb_GetNewAccessToken_Succeed", authWeb.GetNewAccessToken)
}

func (authWeb *AuthWeb) Register(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	defer testWeb.AllSeeder.Down()

	mockAuth := testWeb.AllSeeder.User.UserMock.Data[0]

	bodyRequest := &model_request.RegisterRequest{}
	bodyRequest.Username = null.NewString(mockAuth.Username.String, true)
	bodyRequest.Email = null.NewString(mockAuth.Email.String, true)
	bodyRequest.Password = null.NewString(mockAuth.Password.String, true)

	bodyRequestJsonByte, marshalErr := json.Marshal(bodyRequest)
	if marshalErr != nil {
		t.Fatal(marshalErr)
	}
	bodyRequestBuffer := bytes.NewBuffer(bodyRequestJsonByte)

	url := fmt.Sprintf("%s/%s/register", testWeb.Server.URL, authWeb.Path)
	request, newRequestErr := http.NewRequest(http.MethodPost, url, bodyRequestBuffer)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}

	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	bodyResponse := &model_response.Response[*entity.User]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, mockAuth.Username.String, bodyResponse.Data.Username.String)
	assert.Equal(t, mockAuth.Email.String, bodyResponse.Data.Email.String)
	assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(bodyResponse.Data.Password.String), []byte(mockAuth.Password.String)))

	newUserMock := bodyResponse.Data
	testWeb.AllSeeder.User.UserMock.Data = append(testWeb.AllSeeder.User.UserMock.Data, newUserMock)
}

func (authWeb *AuthWeb) Login(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedUserMock := testWeb.AllSeeder.User.UserMock.Data[0]

	bodyRequest := &model_request.LoginRequest{}
	bodyRequest.Email = selectedUserMock.Email
	bodyRequest.Password = selectedUserMock.Password

	bodyRequestJsonByte, marshalErr := json.Marshal(bodyRequest)
	if marshalErr != nil {
		t.Fatal(marshalErr)
	}
	bodyRequestBuffer := bytes.NewBuffer(bodyRequestJsonByte)

	url := fmt.Sprintf("%s/%s/login", testWeb.Server.URL, authWeb.Path)
	request, newRequestErr := http.NewRequest(http.MethodPost, url, bodyRequestBuffer)
	if newRequestErr != nil {
		t.Fatal(newRequestErr)
	}

	response, doErr := http.DefaultClient.Do(request)
	if doErr != nil {
		t.Fatal(doErr)
	}

	bodyResponse := &model_response.Response[*entity.Session]{}
	decodeErr := json.NewDecoder(response.Body).Decode(bodyResponse)
	if decodeErr != nil {
		t.Fatal(decodeErr)
	}

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, selectedUserMock.Id, bodyResponse.Data.UserId)
}

func (authWeb *AuthWeb) Logout(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()

	selectedSessionMock := testWeb.AllSeeder.Session.SessionMock.Data[0]
	logoutURL := fmt.Sprintf("%s/%s/logout", testWeb.Server.URL, authWeb.Path)
	logoutRequest, newLogoutRequestErr := http.NewRequest(http.MethodPost, logoutURL, nil)
	if newLogoutRequestErr != nil {
		t.Fatal(newLogoutRequestErr)
	}

	accessToken := selectedSessionMock.AccessToken
	logoutRequest.Header.Set("Authorization", "Bearer "+accessToken.String)

	logoutResponse, logoutDoErr := http.DefaultClient.Do(logoutRequest)
	if logoutDoErr != nil {
		t.Fatal(logoutDoErr)
	}
	defer logoutResponse.Body.Close()

	assert.Equal(t, http.StatusOK, logoutResponse.StatusCode)
}
func (authWeb *AuthWeb) GetNewAccessToken(t *testing.T) {
	t.Parallel()

	testWeb := GetTestWeb()
	testWeb.AllSeeder.Up()
	defer testWeb.AllSeeder.Down()

	selectedSessionMock := testWeb.AllSeeder.Session.SessionMock.Data[0]
	newAccessTokenURL := fmt.Sprintf("%s/%s/getNewAccessToken", testWeb.Server.URL, authWeb.Path)
	newAccessTokenRequest, newAccessTokenRequestErr := http.NewRequest(http.MethodPost, newAccessTokenURL, nil)
	if newAccessTokenRequestErr != nil {
		t.Fatal(newAccessTokenRequestErr)
	}

	refreshToken := selectedSessionMock.RefreshToken
	newAccessTokenRequest.Header.Set("Authorization", "Bearer "+refreshToken.String)

	newAccessTokenResponse, newAccessTokenDoErr := http.DefaultClient.Do(newAccessTokenRequest)
	if newAccessTokenDoErr != nil {
		t.Fatal(newAccessTokenDoErr)
	}
	defer newAccessTokenResponse.Body.Close()

	assert.Equal(t, http.StatusOK, newAccessTokenResponse.StatusCode)
}
