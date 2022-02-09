package mailer

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient(t *testing.T) {
	t.Parallel()
	require.NotNil(t, NewClient(&Params{}))
}

func TestMailError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		err    error
		expect error
	}{
		{
			name:   "error is nil",
			err:    nil,
			expect: nil,
		},
		{
			name:   "error is invalid error",
			err:    errors.New("some error"),
			expect: ErrUnknown,
		},
		{
			name:   "context to canceled",
			err:    context.Canceled,
			expect: ErrTimeout,
		},
		{
			name:   "context to deadline exceeded",
			err:    context.DeadlineExceeded,
			expect: ErrTimeout,
		},
		{
			name:   "send grid error to bad request",
			err:    &SendGridError{Code: http.StatusBadRequest},
			expect: ErrInvalidArgument,
		},
		{
			name:   "send grid error to unauthenticated",
			err:    &SendGridError{Code: http.StatusUnauthorized},
			expect: ErrUnauthenticated,
		},
		{
			name:   "send grid error to forbidden",
			err:    &SendGridError{Code: http.StatusForbidden},
			expect: ErrPermissionDenied,
		},
		{
			name:   "send grid error to request entity too large",
			err:    &SendGridError{Code: http.StatusRequestEntityTooLarge},
			expect: ErrPayloadTooLong,
		},
		{
			name:   "send grid error to not found",
			err:    &SendGridError{Code: http.StatusNotFound},
			expect: ErrNotFound,
		},
		{
			name:   "send grid error to internal server error",
			err:    &SendGridError{Code: http.StatusInternalServerError},
			expect: ErrInternal,
		},
		{
			name:   "send grid error to bad gateway",
			err:    &SendGridError{Code: http.StatusBadGateway},
			expect: ErrUnavailable,
		},
		{
			name:   "send grid error to gateway timeout",
			err:    &SendGridError{Code: http.StatusGatewayTimeout},
			expect: ErrTimeout,
		},
		{
			name:   "send grid error to unexpected error",
			err:    &SendGridError{Code: http.StatusNotImplemented},
			expect: ErrUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.ErrorIs(t, mailError(tt.err), tt.expect)
		})
	}
}
