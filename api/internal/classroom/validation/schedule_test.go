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

func TestUpdateSchedules(t *testing.T) {
	t.Parallel()
	validator := NewRequestValidation()

	tests := []struct {
		name  string
		req   *classroom.UpdateSchedulesRequest
		isErr bool
	}{
		{
			name: "success",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  int32(time.Sunday),
						IsClosed: true,
						Lessons:  []*classroom.ScheduleToUpdate_Lesson{},
					},
					{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.ScheduleToUpdate_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
			},
			isErr: false,
		},
		{
			name: "Schedules is min_items",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{},
			},
			isErr: true,
		},
		{
			name: "Schedules.Weekday is gte",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  -1,
						IsClosed: true,
						Lessons:  []*classroom.ScheduleToUpdate_Lesson{},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Schedules.Weekday is lte",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  7,
						IsClosed: true,
						Lessons:  []*classroom.ScheduleToUpdate_Lesson{},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Schedules.Lessons.StartTime is len",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.ScheduleToUpdate_Lesson{
							{StartTime: "", EndTime: "1830"},
						},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Schedules.Lessons.StartTime is pattern",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.ScheduleToUpdate_Lesson{
							{StartTime: "abcd", EndTime: "1830"},
						},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Schedules.Lessons.EndTime is len",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.ScheduleToUpdate_Lesson{
							{StartTime: "1700", EndTime: ""},
						},
					},
				},
			},
			isErr: true,
		},
		{
			name: "Schedules.Lessons.EndTime is pattern",
			req: &classroom.UpdateSchedulesRequest{
				Schedules: []*classroom.ScheduleToUpdate{
					{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.ScheduleToUpdate_Lesson{
							{StartTime: "1700", EndTime: "abcd"},
						},
					},
				},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validator.UpdateSchedules(tt.req)
			assert.Equal(t, tt.isErr, err != nil, err)
		})
	}
}
