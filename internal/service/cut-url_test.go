package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestCutUrl(t *testing.T) {
	tests := []test.Test[string, string]{
		{
			Name: "base case",
			Args: "http://google.com",
			Want: env.URL() + "/11223344",
		},
		{
			Name:      "incorrect URL",
			Args:      "htp:/google.com",
			WantError: true,
		},
		{
			Name:      "empty URL",
			Args:      "",
			WantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			url, err := CutURL(tt.Args)

			if tt.WantError {
				require.Error(t, err)

				return
			}

			require.Len(t, url, len(tt.Want))
			require.NoError(t, err)
		})
	}
}
