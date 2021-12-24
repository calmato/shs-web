package validation

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestListSchedules(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.ListSchedulesRequest
		isErr bool
	}{
		{
			name:  "success",
			req:   &classroom.ListSchedulesRequest{},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.ListSchedules(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}

func TestGetSchedule(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.GetScheduleRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.GetScheduleRequest{
				Weekday: int32(time.Monday),
			},
			isErr: false,
		},
		{
			name: "Weekday is gte",
			req: &classroom.GetScheduleRequest{
				Weekday: -1,
			},
			isErr: true,
		},
		{
			name: "Weekday is lte",
			req: &classroom.GetScheduleRequest{
				Weekday: 7,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.GetSchedule(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
