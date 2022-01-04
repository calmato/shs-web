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

func TestScheduleToUpdate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		schedule *ScheduleToUpdate
		expect   *classroom.ScheduleToUpdate
	}{
		{
			name: "success to closed",
			schedule: &ScheduleToUpdate{
				Weekday:  time.Monday,
				IsClosed: true,
				Lessons: ScheduleLessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
					{StartTime: "2000", EndTime: "2130"},
				},
			},
			expect: &classroom.ScheduleToUpdate{
				Weekday:  int32(time.Monday),
				IsClosed: true,
			},
		},
		{
			name: "success to weekday",
			schedule: &ScheduleToUpdate{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons: ScheduleLessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
					{StartTime: "2000", EndTime: "2130"},
				},
			},
			expect: &classroom.ScheduleToUpdate{
				Weekday:  int32(time.Monday),
				IsClosed: false,
				Lessons: []*classroom.ScheduleToUpdate_Lesson{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
					{StartTime: "2000", EndTime: "2130"},
				},
			},
		},
		{
			name: "success to holiday",
			schedule: &ScheduleToUpdate{
				Weekday:  time.Sunday,
				IsClosed: false,
				Lessons: ScheduleLessons{
					{StartTime: "1530", EndTime: "1700"},
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
					{StartTime: "2000", EndTime: "2130"},
				},
			},
			expect: &classroom.ScheduleToUpdate{
				Weekday:  int32(time.Sunday),
				IsClosed: false,
				Lessons: []*classroom.ScheduleToUpdate_Lesson{
					{StartTime: "1530", EndTime: "1700"},
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "2000"},
					{StartTime: "2000", EndTime: "2130"},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewScheduleToUpdate(tt.schedule))
		})
	}
}

func TestSchedulesToUpdate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules []*ScheduleToUpdate
		expect    []*classroom.ScheduleToUpdate
	}{
		{
			name: "success",
			schedules: []*ScheduleToUpdate{
				{
					Weekday:  time.Sunday,
					IsClosed: false,
					Lessons: ScheduleLessons{
						{StartTime: "1530", EndTime: "1700"},
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: ScheduleLessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
				{
					Weekday:  time.Tuesday,
					IsClosed: true,
					Lessons: ScheduleLessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
			},
			expect: []*classroom.ScheduleToUpdate{
				{
					Weekday:  int32(time.Sunday),
					IsClosed: false,
					Lessons: []*classroom.ScheduleToUpdate_Lesson{
						{StartTime: "1530", EndTime: "1700"},
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
				{
					Weekday:  int32(time.Monday),
					IsClosed: false,
					Lessons: []*classroom.ScheduleToUpdate_Lesson{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "2000"},
						{StartTime: "2000", EndTime: "2130"},
					},
				},
				{
					Weekday:  int32(time.Tuesday),
					IsClosed: true,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSchedulesToUpdate(tt.schedules))
		})
	}
}
