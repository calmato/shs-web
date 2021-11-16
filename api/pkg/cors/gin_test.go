package cors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGinMiddleware(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, NewGinMiddleware())
}
