package main

import (
	"app/internal/server"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}
}

func main() {
	server.RunServer()
}
