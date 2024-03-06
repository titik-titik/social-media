package main

import (
	"fmt"
	"log"
	"net/http"
	"social-media/container"
	"social-media/db/redis"
)

func main() {
	fmt.Println("Web started.")

	webContainer := container.NewWebContainer()

	address := fmt.Sprintf(
		"%s:%s",
		webContainer.Env.App.Host,
		webContainer.Env.App.Port,
	)
	redisManager, err := redis.NewRedisConnection()
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	defer func() {
		if err := redisManager.Close(); err != nil {
			log.Printf("Failed to close Redis connection: %v", err)
		}
	}()
	listenAndServeErr := http.ListenAndServe(address, webContainer.Route.Router)
	if listenAndServeErr != nil {
		panic(listenAndServeErr)
	}
	fmt.Println("Web finished.")
}
