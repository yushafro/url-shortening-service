package handler

import (
	"fmt"
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/service"
)

func GetUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	originUrl := service.Urls[id]

	if originUrl == "" {
		err := fmt.Sprintf(NO_SUCH_ITEM_BY_ID, originUrl, id)
		http.Error(w, err, http.StatusBadRequest)

		return
	}

	w.Header().Add(LOCATION, originUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
