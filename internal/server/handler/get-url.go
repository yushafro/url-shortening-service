package handler

import (
	"fmt"
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/service"
)

func GetURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	originURL := service.Urls[id]

	if originURL == "" {
		err := fmt.Sprintf(NoSuchItemByID, originURL, id)
		http.Error(w, err, http.StatusBadRequest)

		return
	}

	w.Header().Add(Location, originURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
