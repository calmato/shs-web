package grpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	tests := []struct {
		name   string
		server *grpc.Server
		port   int64
		expect Server
		isErr  bool
	}{
		{
			name:   "success",
			server: &grpc.Server{},
			port:   28080,
			isErr:  false,
		},
		{
			name:   "failed",
			server: &grpc.Server{},
			port:   -1,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewGRPCServer(tt.server, tt.port)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, actual)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, actual)
		})
	}
}
