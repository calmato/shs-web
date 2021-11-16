package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGinMiddleware(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		params *Params
		isErr  bool
	}{
		{
			name: "success",
			params: &Params{
				Level: "debug",
				Path:  "",
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			logger, err := NewGinMiddleware(tt.params)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, logger)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, logger)
		})
	}
}
