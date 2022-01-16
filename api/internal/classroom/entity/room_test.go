package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestRoom_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		room   *Room
		expect *classroom.Room
	}{
		{
			name: "success",
			room: &Room{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expect: &classroom.Room{
				Id:        1,
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.room.Proto())
		})
	}
}

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
