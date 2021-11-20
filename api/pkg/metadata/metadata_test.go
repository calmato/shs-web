package metadata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestGet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		ctx    context.Context
		key    string
		expect string
		isErr  bool
	}{
		{
			name:   "success",
			ctx:    metadata.NewIncomingContext(context.Background(), metadata.Pairs("debug", "test")),
			key:    "debug",
			expect: "test",
			isErr:  false,
		},
		{
			name:   "invalid metadata",
			ctx:    context.Background(),
			key:    "",
			expect: "",
			isErr:  true,
		},
		{
			name:   "not found metadata",
			ctx:    metadata.NewIncomingContext(context.Background(), metadata.MD{}),
			key:    "debug",
			expect: "",
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := Get(tt.ctx, tt.key)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestSet(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		ctx       context.Context
		key       string
		value     string
		isSuccess bool
	}{
		{
			name:      "success",
			ctx:       metadata.NewOutgoingContext(context.Background(), metadata.MD{}),
			key:       "debug",
			value:     "test",
			isSuccess: true,
		},
		{
			name:      "failed to set metadata",
			ctx:       metadata.NewOutgoingContext(context.Background(), metadata.MD{}),
			key:       "",
			value:     "",
			isSuccess: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := Set(tt.ctx, tt.key, tt.value)
			md, ok := metadata.FromOutgoingContext(ctx)
			if ok {
				v := md.Get(tt.key)
				assert.Contains(t, v, tt.value)
			} else {
				t.FailNow()
			}
		})
	}
}
