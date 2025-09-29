package url_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/internal/config/env"
	"github.com/yushafro/url-shortening-service/internal/url"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestCutUrl(t *testing.T) {
	t.Parallel()

	tests := []test.Test[string, string]{
		{
			Name: "base case",
			Args: "http://google.com",
			Want: env.URL() + "/11223344",
		},
		{
			Name:  "incorrect URL",
			Args:  "htp:/google.com",
			Error: url.ErrInvalidURL,
		},
		{
			Name:  "empty URL",
			Args:  "",
			Error: url.ErrInvalidURL,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			url, err := url.CutURL(test.Args)

			if test.Error != nil {
				require.EqualError(t, err, test.Error.Error())

				return
			}

			require.NoError(t, err)
			require.Len(t, url, len(test.Want))
		})
	}
}
