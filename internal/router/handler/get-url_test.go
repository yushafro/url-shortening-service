package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetURL(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc(GetURLPattern, GetURLHandler)
	mux.HandleFunc(CutURLPattern, CutURLHandler)

	server := httptest.NewServer(mux)
	defer server.Close()

	tests := []test{
		{
			Name: "valid URL",
			Args: args{
				Method: http.MethodGet,
				Body:   "https://google.com",
			},
			Want: want{
				StatusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, id := postID(testServer{
				t:      t,
				server: server,
				test:   tt,
			})

			checkValidID(testServer{
				t:      t,
				server: server,
				test:   tt,
			}, id)

			checkInvalidID(testServer{
				t:      t,
				server: server,
				test:   tt,
			})
		})
	}
}

func checkValidID(ts testServer, id string) {
	req, _ := http.NewRequest(ts.test.Args.Method, ts.server.URL+"/"+id, nil)
	res, err := ts.server.Client().Do(req)

	require.NoError(ts.t, err)
	defer res.Body.Close()

	require.Equal(ts.t, ts.test.Want.StatusCode, res.StatusCode)
}

func checkInvalidID(ts testServer) {
	req, _ := http.NewRequest(ts.test.Args.Method, ts.server.URL+"/"+"invalidID", nil)
	res, err := ts.server.Client().Do(req)

	require.NoError(ts.t, err)
	defer res.Body.Close()

	require.Equal(ts.t, http.StatusBadRequest, res.StatusCode)
}
