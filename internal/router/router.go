package router

import (
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/router/handler"
)

func InitServer() {
	mux := http.NewServeMux()

	mux.HandleFunc(handler.GetURLPattern, handler.GetURLHandler)
	mux.HandleFunc(handler.CutURLPattern, handler.CutURLHandler)

	err := http.ListenAndServe(":"+env.Port(), mux)
	if err != nil {
		panic(err)
	}
}
