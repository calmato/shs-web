package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		schedule *classroom.Schedule
		expect   *Schedule
	}{
		{
			name: "success",
			schedule: &classroom.Schedule{
				Weekday:  int32(time.Monday),
				IsClosed: false,
				Lessons: []*classroom.Schedule_Lesson{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
					{StartTime: "2000", EndTime: "2130"},
				},
			},
			expect: &Schedule{
				Schedule: &classroom.Schedule{
					Weekday:  int32(time.Monday),
					IsClosed: false,
					Lessons: []*classroom.Schedule_Lesson{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchedule(tt.schedule))
		})
	}
}

func TestSchedules(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules []*classroom.Schedule
		expect    Schedules
	}{
		{
			name: "success",
			schedules: []*classroom.Schedule{
				{
					Weekday:  int32(time.Sunday),
					IsClosed: true,
					Lessons:  []*classroom.Schedule_Lesson{},
				},
				{
					Weekday:  int32(time.Monday),
					IsClosed: false,
					Lessons: []*classroom.Schedule_Lesson{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
			},
			expect: Schedules{
				{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Sunday),
						IsClosed: true,
						Lessons:  []*classroom.Schedule_Lesson{},
					},
				},
				{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchedules(tt.schedules))
		})
	}
}

func TestSchedules_Map(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules Schedules
		expect    map[time.Weekday]*Schedule
	}{
		{
			name: "success",
			schedules: Schedules{
				{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Sunday),
						IsClosed: true,
						Lessons:  []*classroom.Schedule_Lesson{},
					},
				},
				{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
				},
			},
			expect: map[time.Weekday]*Schedule{
				time.Sunday: {
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Sunday),
						IsClosed: true,
						Lessons:  []*classroom.Schedule_Lesson{},
					},
				},
				time.Monday: {
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
							{StartTime: "2000", EndTime: "2130"},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.schedules.Map())
		})
	}
}
