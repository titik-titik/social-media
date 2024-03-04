package main

import (
	"fmt"
	"net/http"
	"social-media/container"
)

func main() {
	fmt.Println("Web started.")

	WebContainer := container.NewWebContainer()
	// Setup Engine
	address := fmt.Sprintf(
		"%s:%s",
		WebContainer.Env.App.Host,
		WebContainer.Env.App.Port,
	)
	listenAndServeErr := http.ListenAndServe(address, WebContainer.Route.Router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}

	fmt.Println("Web finished.")
}
