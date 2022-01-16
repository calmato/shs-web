package validation

import (
	"testing"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestGetRoom(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.GetRoomRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.GetRoomRequest{
				Id: 1,
			},
			isErr: false,
		},
		{
			name: "Id is gt",
			req: &classroom.GetRoomRequest{
				Id: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetRoom(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetRoomsTotal(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.GetRoomsTotalRequest
		isErr bool
	}{
		{
			name:  "success",
			req:   &classroom.GetRoomsTotalRequest{},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetRoomsTotal(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestUpdateRoomsTotal(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.UpdateRoomsTotalRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.UpdateRoomsTotalRequest{
				Total: 1,
			},
			isErr: false,
		},
		{
			name: "Total is gt",
			req: &classroom.UpdateRoomsTotalRequest{
				Total: 0,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateRoomsTotal(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
