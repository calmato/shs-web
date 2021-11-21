package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrpcServerOptions(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		params *OptionParams
	}{
		{
			name:   "success",
			params: &OptionParams{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.NotNil(t, NewGRPCOptions(tt.params))
		})
	}
}
