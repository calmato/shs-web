package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestScheduleToUpdate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules []*classroom.ScheduleToUpdate
		expect    Schedules
	}{
		{
			name: "success",
			schedules: []*classroom.ScheduleToUpdate{
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
						{StartTime: "1830", EndTime: "1900"},
					},
				},
			},
			expect: Schedules{
				{
					Weekday:  time.Sunday,
					IsClosed: true,
					Lessons:  Lessons{},
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
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

func TestSchedule_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		schedule *Schedule
		expect   *Schedule
		isErr    bool
	}{
		{
			name: "success to closed school",
			schedule: &Schedule{
				Weekday:     time.Monday,
				IsClosed:    true,
				LessonsJSON: datatypes.JSON([]byte(`null`)),
			},
			expect: &Schedule{
				Weekday:     time.Monday,
				IsClosed:    true,
				Lessons:     nil,
				LessonsJSON: datatypes.JSON([]byte(`null`)),
			},
			isErr: false,
		},
		{
			name: "success to not closed school",
			schedule: &Schedule{
				Weekday:     time.Monday,
				IsClosed:    false,
				LessonsJSON: datatypes.JSON([]byte(`[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"1900"}]`)),
			},
			expect: &Schedule{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons: Lessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "1900"},
				},
				LessonsJSON: datatypes.JSON([]byte(`[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"1900"}]`)),
			},
			isErr: false,
		},
		{
			name: "failed to unmarshal lessons json",
			schedule: &Schedule{
				Weekday:     time.Monday,
				IsClosed:    false,
				LessonsJSON: nil,
			},
			expect: &Schedule{
				Weekday:     time.Monday,
				IsClosed:    false,
				Lessons:     nil,
				LessonsJSON: nil,
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.schedule.Fill()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.schedule)
		})
	}
}

func TestSchedule_FillJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		schedule *Schedule
		expect   *Schedule
		isErr    bool
	}{
		{
			name: "success to lessons is nil",
			schedule: &Schedule{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons:  nil,
			},
			expect: &Schedule{
				Weekday:     time.Monday,
				IsClosed:    false,
				Lessons:     nil,
				LessonsJSON: datatypes.JSON([]byte(`null`)),
			},
			isErr: false,
		},
		{
			name: "success to lessons is not nil",
			schedule: &Schedule{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons: Lessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "1900"},
				},
			},
			expect: &Schedule{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons: Lessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "1900"},
				},
				LessonsJSON: datatypes.JSON([]byte(`[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"1900"}]`)),
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.schedule.FillJSON()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.schedule)
		})
	}
}

func TestSchedule_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name     string
		schedule *Schedule
		expect   *classroom.Schedule
	}{
		{
			name: "success",
			schedule: &Schedule{
				Weekday:  time.Monday,
				IsClosed: false,
				Lessons: Lessons{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "1900"},
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expect: &classroom.Schedule{
				Weekday:  int32(time.Monday),
				IsClosed: false,
				Lessons: []*classroom.Schedule_Lesson{
					{StartTime: "1700", EndTime: "1830"},
					{StartTime: "1830", EndTime: "1900"},
				},
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.schedule.Proto())
		})
	}
}

func TestSchedules_Weekdays(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules Schedules
		expect    []time.Weekday
	}{
		{
			name: "success",
			schedules: Schedules{
				{
					Weekday:  time.Sunday,
					IsClosed: true,
					Lessons:  Lessons{},
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
				},
			},
			expect: []time.Weekday{
				time.Sunday,
				time.Monday,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.schedules.Weekdays())
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
					Weekday:  time.Sunday,
					IsClosed: true,
					Lessons:  Lessons{},
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
				},
			},
			expect: map[time.Weekday]*Schedule{
				time.Sunday: {
					Weekday:  time.Sunday,
					IsClosed: true,
					Lessons:  Lessons{},
				},
				time.Monday: {
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
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

func TestSchedules_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules Schedules
		expect    Schedules
		isErr     bool
	}{
		{
			name: "success to closed school",
			schedules: Schedules{
				{
					Weekday:     time.Sunday,
					IsClosed:    true,
					LessonsJSON: datatypes.JSON([]byte(`null`)),
				},
				{
					Weekday:     time.Monday,
					IsClosed:    false,
					LessonsJSON: datatypes.JSON([]byte(`[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"1900"}]`)),
				},
			},
			expect: Schedules{
				{
					Weekday:     time.Sunday,
					IsClosed:    true,
					Lessons:     nil,
					LessonsJSON: datatypes.JSON([]byte(`null`)),
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
					LessonsJSON: datatypes.JSON([]byte(`[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"1900"}]`)),
				},
			},
			isErr: false,
		},
		{
			name: "failed to unmarshal lessons json",
			schedules: Schedules{
				{
					Weekday:     time.Monday,
					IsClosed:    false,
					LessonsJSON: nil,
				},
			},
			expect: Schedules{
				{
					Weekday:     time.Monday,
					IsClosed:    false,
					Lessons:     nil,
					LessonsJSON: nil,
				},
			},
			isErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.schedules.Fill()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.schedules)
		})
	}
}

func TestSchedules_FillJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules Schedules
		expect    Schedules
		isErr     bool
	}{
		{
			name: "success to lessons is nil",
			schedules: Schedules{
				{
					Weekday:  time.Sunday,
					IsClosed: true,
					Lessons:  nil,
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
				},
			},
			expect: Schedules{
				{
					Weekday:     time.Sunday,
					IsClosed:    true,
					Lessons:     nil,
					LessonsJSON: datatypes.JSON([]byte(`null`)),
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
					LessonsJSON: datatypes.JSON([]byte(`[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"1900"}]`)),
				},
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.schedules.FillJSON()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.schedules)
		})
	}
}

func TestSchedules_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name      string
		schedules Schedules
		expect    []*classroom.Schedule
	}{
		{
			name: "success",
			schedules: Schedules{
				{
					Weekday:   time.Sunday,
					IsClosed:  true,
					Lessons:   nil,
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					Weekday:  time.Monday,
					IsClosed: false,
					Lessons: Lessons{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []*classroom.Schedule{
				{
					Weekday:   int32(time.Sunday),
					IsClosed:  true,
					Lessons:   []*classroom.Schedule_Lesson{},
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				},
				{
					Weekday:  int32(time.Monday),
					IsClosed: false,
					Lessons: []*classroom.Schedule_Lesson{
						{StartTime: "1700", EndTime: "1830"},
						{StartTime: "1830", EndTime: "1900"},
					},
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.schedules.Proto())
		})
	}
}
