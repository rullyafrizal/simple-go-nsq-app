package main

import (
	"log"
	"simple-go-nsq-app/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	handlers := handlers.NewNsqHandler()

	handlers.StartProducing()
	handlers.StartConsuming()
}
