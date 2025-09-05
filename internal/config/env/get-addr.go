package env

import (
	"os"

	"github.com/joho/godotenv"
)

func GetAddr() string {
	godotenv.Load()

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return host + ":" + port
}
