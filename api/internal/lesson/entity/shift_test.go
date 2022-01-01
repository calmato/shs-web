package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		date      time.Time
		startTime string
		endTime   string
		expect    *Shift
	}{
		{
			name:      "success",
			date:      jst.Date(2021, 12, 26, 0, 0, 0, 0),
			startTime: "1700",
			endTime:   "1830",
			expect: &Shift{
				Date:      jst.Date(2021, 12, 26, 0, 0, 0, 0),
				StartTime: "1700",
				EndTime:   "1830",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewShift(tt.date, tt.startTime, tt.endTime)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestShifts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		summary     *ShiftSummary
		schedules   map[time.Weekday]*Schedule
		closedDates []string
		expect      Shifts
		isErr       bool
	}{
		{
			name: "success",
			summary: &ShiftSummary{
				YearMonth: 202202,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
			},
			schedules: map[time.Weekday]*Schedule{
				time.Sunday: {Schedule: &classroom.Schedule{IsClosed: true}},
				time.Monday: {
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1530", EndTime: "1700"},
						},
					},
				},
				time.Tuesday: {
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Tuesday),
						IsClosed: false,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				time.Wednesday: {Schedule: &classroom.Schedule{IsClosed: true}},
				time.Thursday:  {Schedule: &classroom.Schedule{IsClosed: true}},
				time.Friday:    {Schedule: &classroom.Schedule{IsClosed: true}},
				time.Saturday:  {Schedule: &classroom.Schedule{IsClosed: true}},
			},
			closedDates: []string{
				"20220208", "20220221", "20220222",
			},
			expect: Shifts{
				{Date: jst.Date(2022, 2, 1, 0, 0, 0, 0), StartTime: "1700", EndTime: "1830"},
				{Date: jst.Date(2022, 2, 1, 0, 0, 0, 0), StartTime: "1830", EndTime: "2000"},
				{Date: jst.Date(2022, 2, 7, 0, 0, 0, 0), StartTime: "1530", EndTime: "1700"},
				{Date: jst.Date(2022, 2, 14, 0, 0, 0, 0), StartTime: "1530", EndTime: "1700"},
				{Date: jst.Date(2022, 2, 15, 0, 0, 0, 0), StartTime: "1700", EndTime: "1830"},
				{Date: jst.Date(2022, 2, 15, 0, 0, 0, 0), StartTime: "1830", EndTime: "2000"},
				{Date: jst.Date(2022, 2, 28, 0, 0, 0, 0), StartTime: "1530", EndTime: "1700"},
			},
			isErr: false,
		},
		{
			name: "failed to new first and final of month",
			summary: &ShiftSummary{
				YearMonth: 202200,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
			},
			schedules:   map[time.Weekday]*Schedule{},
			closedDates: []string{},
			expect:      Shifts{},
			isErr:       true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := NewShifts(tt.summary, tt.schedules, tt.closedDates)
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Len(t, actual, len(tt.expect))
			for i := range tt.expect {
				assert.Contains(t, actual, tt.expect[i])
			}
		})
	}
}

func TestShift_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shift  *Shift
		expect *lesson.Shift
	}{
		{
			name: "success",
			shift: &Shift{
				ID:             1,
				ShiftSummaryID: 1,
				Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
				StartTime:      "1700",
				EndTime:        "1830",
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expect: &lesson.Shift{
				Id:             1,
				ShiftSummaryId: 1,
				Date:           "20211226",
				StartTime:      "1700",
				EndTime:        "1830",
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.shift.Proto())
		})
	}
}

func TestShifts_GroupByShiftSummaryID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts Shifts
		expect map[int64]Shifts
	}{
		{
			name: "success",
			shifts: Shifts{
				{
					ID:             1,
					ShiftSummaryID: 1,
					Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
				{
					ID:             2,
					ShiftSummaryID: 1,
					Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
				{
					ID:             3,
					ShiftSummaryID: 2,
					Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: map[int64]Shifts{
				1: {
					{
						ID:             1,
						ShiftSummaryID: 1,
						Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now,
						UpdatedAt:      now,
					},
					{
						ID:             2,
						ShiftSummaryID: 1,
						Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now,
						UpdatedAt:      now,
					},
				},
				2: {
					{
						ID:             3,
						ShiftSummaryID: 2,
						Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now,
						UpdatedAt:      now,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.shifts.GroupByShiftSummaryID())
		})
	}
}

func TestShifts_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts Shifts
		expect []*lesson.Shift
	}{
		{
			name: "success",
			shifts: Shifts{
				{
					ID:             1,
					ShiftSummaryID: 1,
					Date:           jst.Date(2021, 12, 26, 0, 0, 0, 0),
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []*lesson.Shift{
				{
					Id:             1,
					ShiftSummaryId: 1,
					Date:           "20211226",
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.shifts.Proto())
		})
	}
}
