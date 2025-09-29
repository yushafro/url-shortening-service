package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/url"
)

const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = 20 * time.Second
	idleTimeout       = 120 * time.Second
)

func InitServer() {
	router := gin.Default()
	router.GET(url.GetURLPattern, url.GetURLHandler)
	router.POST(url.CutURLPattern, url.CutURLHandler)

	server := http.Server{
		Addr:              ":" + env.Port(),
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
