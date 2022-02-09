package pubsub

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const testTopic = "topic-test"

func TestClient(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		projectID string
		isErr     bool
	}{
		{
			name:      "success",
			projectID: "project-test",
			isErr:     false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			client, err := NewClient(ctx, tt.projectID)
			if tt.isErr {
				assert.Error(t, err)
				assert.Nil(t, client)
				return
			}
			assert.NoError(t, err)
			assert.NotNil(t, client)
		})
	}
}

func TestIsRetryable(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		err    error
		expect bool
	}{
		{
			name:   "error is nil",
			err:    nil,
			expect: false,
		},
		{
			name:   "context is deadline exceeded",
			err:    context.DeadlineExceeded,
			expect: true,
		},
		{
			name:   "gRPC error status is internal",
			err:    status.Error(codes.Internal, "some error"),
			expect: true,
		},
		{
			name:   "gRPC error status is unavailable",
			err:    status.Error(codes.Unavailable, "some error"),
			expect: true,
		},
		{
			name:   "gRPC error status is other",
			err:    status.Error(codes.InvalidArgument, "some error"),
			expect: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, isRetryable(tt.err), tt.err)
		})
	}
}

func setEnv() {
	if os.Getenv("PUBSUB_EMULATOR_HOST") == "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", "127.0.0.1:8085")
	}
}
