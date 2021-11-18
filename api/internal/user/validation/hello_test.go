package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.HelloRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.HelloRequest{
				Name: "test",
			},
			isErr: false,
		},
		{
			name: "failed to name is min len",
			req: &user.HelloRequest{
				Name: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.Hello(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
