package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		schedule *entity.Schedule
		expect   *Schedule
	}{
		{
			name: "success",
			schedule: &entity.Schedule{
				Schedule: &classroom.Schedule{
					Weekday:  int32(time.Monday),
					IsClosed: false,
					Lessons: []*classroom.Schedule_Lesson{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
					},
				},
			},
			expect: &Schedule{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons: ScheduleLessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
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
		schedules entity.Schedules
		expect    Schedules
	}{
		{
			name: "success",
			schedules: entity.Schedules{
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
						},
					},
				},
			},
			expect: Schedules{
				{
					Weekday:  time.Sunday,
					IsClosed: true,
					Lessons:  ScheduleLessons{},
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: ScheduleLessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
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
