package web

import (
	"testing"
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
