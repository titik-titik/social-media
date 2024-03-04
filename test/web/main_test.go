package web

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Chdir("../../.")
	fmt.Println("TestWeb started.")
	userWeb := NewUserWeb(t)
	userWeb.Start()
}
