package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/yushafro/url-shortening-service/internal/service"
	pkgHttp "github.com/yushafro/url-shortening-service/pkg/http"
)

const CutURLPattern = "POST /"

var ErrRequiredURL = fmt.Errorf(pkgHttp.Required, "URL")

func CutURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := fmt.Sprintf(pkgHttp.MethodNotAllowed, r.Method)
		http.Error(w, err, http.StatusMethodNotAllowed)

		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	} else if len(body) == 0 {
		http.Error(w, ErrRequiredURL.Error(), http.StatusBadRequest)

		return
	}

	urlID, err := service.CutURL(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(urlID))
}
