package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()
	assert.NotNil(t, New())
}

func TestBase58Encode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		uuid   string
		expect string
	}{
		{
			name:   "success",
			uuid:   "A0EEBC99-9C0B-4EF8-BB6D-6BB9BD380A11",
			expect: "kSByoE6FetnPs5Byk3a9Zx",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, Base58Encode(tt.uuid))
		})
	}
}
