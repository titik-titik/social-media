package test

import (
	"fmt"
	"os"
	"social-media/test/web"
	"testing"
)

func Test(t *testing.T) {
	chdirErr := os.Chdir("../.")
	if chdirErr != nil {
		t.Fatal(chdirErr)
	}
	fmt.Println("TestWeb started.")
	authWeb := web.NewAuthWeb(t)
	authWeb.Start()
	userWeb := web.NewUserWeb(t)
	userWeb.Start()
	postWeb := web.NewPostWeb(t)
	postWeb.Start()
	fmt.Println("TestWeb finished.")
}
