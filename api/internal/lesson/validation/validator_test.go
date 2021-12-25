package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestValidation(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		expect RequestValidation
	}{
		{
			name:   "success",
			expect: &requestValidation{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewRequestValidation())
		})
	}
}

func TestValidationError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		message string
		expect  error
	}{
		{
			name:    "success",
			message: "test message",
			expect:  fmt.Errorf("%w: %s", ErrRequestValidation, "test message"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, validationError(tt.message))
		})
	}
}
