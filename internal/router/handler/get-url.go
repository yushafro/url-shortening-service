package handler

import (
	"fmt"
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/service"
	pkgHttp "github.com/yushafro/url-shortening-service/pkg/http"
)

const GetURLPattern = "GET /{id}"

func GetURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	originURL := service.URLs[id]

	if originURL == "" {
		err := fmt.Sprintf(pkgHttp.NoFoundByID, "url", id)
		http.Error(w, err, http.StatusBadRequest)

		return
	}

	w.Header().Add(pkgHttp.Location, originURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
