package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRooms(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		size   int
		expect Rooms
	}{
		{
			name: "success",
			size: 3,
			expect: Rooms{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewRooms(tt.size))
		})
	}
}
