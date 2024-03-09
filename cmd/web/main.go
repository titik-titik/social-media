package main

import (
	"fmt"
	"net/http"
	"social-media/container"
)

func main() {
	fmt.Println("Web started.")

	webContainer := container.NewWebContainer()

	address := fmt.Sprintf(
		"%s:%s",
		"0.0.0.0",
		webContainer.Env.App.Port,
	)
	listenAndServeErr := http.ListenAndServe(address, webContainer.Route.Router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}
	fmt.Println("Web finished.")
}
