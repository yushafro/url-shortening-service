package id_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/pkg/id"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestRandomID(t *testing.T) {
	t.Parallel()

	tests := []test.Test[uint8, int]{
		{
			Name: "small ID",
			Args: 8,
			Want: 8,
		},
		{
			Name: "big ID",
			Args: 200,
			Want: 200,
		},
		{
			Name:  "empty length",
			Args:  0,
			Error: id.ErrEmptyID,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			randomID, err := id.RandomID(test.Args)

			if test.Error != nil {
				require.EqualError(t, err, test.Error.Error())

				return
			}

			require.NoError(t, err)
			require.Len(t, randomID, test.Want)
		})
	}
}
