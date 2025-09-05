package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/service"
)

func CutUrlHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	} else if len(body) == 0 {
		err := fmt.Sprintf(REQUIRED, "URL")
		http.Error(w, err, http.StatusBadRequest)

		return
	}

	urlId := service.CutUrl(string(body))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(urlId))
}
