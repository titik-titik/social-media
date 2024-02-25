package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"social-media/internal/config"
)

func main() {
	fmt.Println("Web starting.")

	errEnvLoad := godotenv.Load()
	if errEnvLoad != nil {
		fmt.Println(fmt.Sprintf("Error loading .env file: %+v", errEnvLoad))
	}

	envConfig := config.NewEnvConfig()
	bootstrapConfig := &config.BootstrapConfig{
		Env: envConfig,
	}
	config.Bootstrap(bootstrapConfig)

	fmt.Println("Web started.")
}
