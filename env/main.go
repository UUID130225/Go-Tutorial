package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load eviroment:", err)
	}

	return nil
}

func main() {
	loadENV()

	appEnv := os.Getenv("APP_ENV")
	fmt.Println(appEnv)
}