package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/service"
)

func TestCutUrl(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(CutURLPattern, CutURLHandler)

	server := httptest.NewServer(mux)
	defer server.Close()

	tests := []test{
		{
			Name: "valid request",
			Args: args{
				Method: http.MethodPost,
				Path:   "/",
				Body:   "https://google.com",
			},
			Want: want{
				StatusCode: http.StatusCreated,
				Response:   env.URL() + "/11223344",
			},
		},
		{
			Name: "empty body (URL)",
			Args: args{
				Method: http.MethodPost,
				Path:   "/",
				Body:   "",
			},
			Want: want{
				StatusCode: http.StatusBadRequest,
				Response:   ErrRequiredURL.Error() + "\n",
			},
			WantError: true,
		},
		{
			Name: "invalid body (URL)",
			Args: args{
				Method: http.MethodPost,
				Path:   "/",
				Body:   "htp:/google.com",
			},
			Want: want{
				StatusCode: http.StatusBadRequest,
				Response:   service.ErrInvalidURL.Error() + "\n",
			},
			WantError: true,
		},
		// {
		// 	Name: "GET request",
		// 	Args: args{
		// 		Method: http.MethodGet,
		// 		Path:   "/",
		// 		Body:   "https://google.com",
		// 	},
		// 	Want: want{
		// 		StatusCode: http.StatusMethodNotAllowed,
		// 		Response:   fmt.Sprintf(pkgHttp.MethodNotAllowed, http.MethodGet) + "\n",
		// 	},
		// 	WantError: true,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			body, _ := postID(testServer{
				t:      t,
				server: server,
				test:   tt,
			})

			if tt.WantError {
				require.Equal(t, tt.Want.Response, string(body))

				return
			}

			require.Len(t, string(body), len(tt.Want.Response))
		})
	}
}

func postID(ts testServer) ([]byte, string) {
	req, _ := http.NewRequest(http.MethodPost, ts.server.URL, strings.NewReader(ts.test.Args.Body))
	res, err := ts.server.Client().Do(req)
	require.NoError(ts.t, err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	require.NoError(ts.t, err)

	return body, string(body[len(body)-8:])
}
