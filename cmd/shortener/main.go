package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/yushafro/url-shortening-service/internal/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("cmd/shortener/main: failed to load env")
	}

	router.InitServer()
}
