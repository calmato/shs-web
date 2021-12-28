package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestGetStudent(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *user.GetStudentRequest
		isErr bool
	}{
		{
			name: "success",
			req: &user.GetStudentRequest{
				Id: "MvcNyJFfdoDTrqC1KDHbRw",
			},
			isErr: false,
		},
		{
			name: "Id is min_len",
			req: &user.GetStudentRequest{
				Id: "",
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetStudent(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
