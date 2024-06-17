package main

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv() Env {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Env{
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Port:        port,
	}
}
