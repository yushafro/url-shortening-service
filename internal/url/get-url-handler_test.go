package url_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/url"
	"github.com/yushafro/url-shortening-service/pkg/httputils"
)

func TestGetURL(t *testing.T) {
	t.Parallel()

	tests := []tableTest{
		{
			Name: "valid ID",
			Args: request{
				Method:      http.MethodGet,
				Body:        "url=https://google.com",
				ContentType: httputils.URLEncoded,
			},
			Want: response{
				StatusCode: http.StatusOK,
			},
		},
		{
			Name: "invalid ID",
			Args: request{
				Method:      http.MethodGet,
				Body:        "url=https://google.com",
				ContentType: httputils.URLEncoded,
			},
			Want: response{
				StatusCode: http.StatusBadRequest,
				WantError:  true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			router := gin.New()
			router.GET(url.GetURLPattern, url.GetURLHandler)
			router.POST(url.CutURLPattern, url.CutURLHandler)

			server := httptest.NewServer(router)
			defer server.Close()

			testServer := &TestServer{
				T:      t,
				Server: server,
				Test:   test,
			}

			body := postID(testServer)
			urlID := body[len(body)-8:]

			if test.Want.WantError {
				checkID(testServer, "invalidID")

				return
			}

			checkID(testServer, urlID)
		})
	}
}

func checkID(testServer *TestServer, urlID string) {
	res := req.DevMode().Get(testServer.Server.URL + "/" + urlID).Do()

	require.NoError(testServer.T, res.Err)
	require.Equal(testServer.T, testServer.Test.Want.StatusCode, res.StatusCode)
}
