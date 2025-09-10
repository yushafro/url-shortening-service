package id

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yushafro/url-shortening-service/pkg/test"
)

func TestRandomID(t *testing.T) {
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
			Name:      "empty length",
			Args:      0,
			WantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			id, err := RandomID(tt.Args)

			if tt.WantError {
				require.Error(t, err)

				return
			}

			require.Equal(t, tt.Want, len(id))
			require.NoError(t, err)
		})
	}
}
