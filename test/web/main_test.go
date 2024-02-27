package web

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("TestWeb started.")
	userWeb := NewUserWeb(t)
	userWeb.Start()
}
