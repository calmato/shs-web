package mailer

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendGridError_Error(t *testing.T) {
	t.Parallel()
	err := &SendGridError{
		Code:    http.StatusBadRequest,
		Message: "test error",
		Field:   "field",
		Help:    "help message",
	}
	assert.Equal(t, "mail: failed to SendGrid. code=400, field=test error", err.Error())
}

func TestSubstitutions(t *testing.T) {
	t.Parallel()
	params := map[string]string{"key": "value"}
	expect := map[string]interface{}{"key": "value"}
	assert.Equal(t, expect, NewSubstitutions(params))
}
