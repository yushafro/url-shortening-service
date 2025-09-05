package server

import (
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/server/handler"
)

func InitServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{id}", handler.GetURLHandler)
	mux.HandleFunc("POST /", handler.CutURLHandler)

	err := http.ListenAndServe(env.GetAddr(), mux)
	if err != nil {
		panic(err)
	}
}
