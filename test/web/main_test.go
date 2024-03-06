package web

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	chdirErr := os.Chdir("../../.")
	if chdirErr != nil {
		t.Fatal(chdirErr)
	}
	fmt.Println("TestWeb started.")
	authWeb := NewAuthWeb(t)
	authWeb.Start()
	userWeb := NewUserWeb(t)
	userWeb.Start()
}
