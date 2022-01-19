package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentShift(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name    string
		shift   *entity.Shift
		enabled bool
		expect  *StudentShift
	}{
		{
			name: "success",
			shift: &entity.Shift{
				Shift: &lesson.Shift{
					Id:             1,
					ShiftSummaryId: 1,
					Date:           "20211226",
					StartTime:      "1700",
					EndTime:        "1830",
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			enabled: true,
			expect: &StudentShift{
				ID:        1,
				Enabled:   true,
				StartTime: "1700",
				EndTime:   "1830",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentShift(tt.shift, tt.enabled))
		})
	}
}

func TestStudentShifts(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name          string
		shifts        entity.Shifts
		studentShifts map[int64]*entity.StudentShift
		onlyEnabled   bool
		expect        StudentShifts
	}{
		{
			name: "success",
			shifts: entity.Shifts{
				{
					Shift: &lesson.Shift{
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
			studentShifts: map[int64]*entity.StudentShift{},
			onlyEnabled:   false,
			expect: StudentShifts{
				{
					ID:        1,
					Enabled:   false,
					StartTime: "1700",
					EndTime:   "1830",
				},
			},
		},
		{
			name: "success to only enabled",
			shifts: entity.Shifts{
				{
					Shift: &lesson.Shift{
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
			studentShifts: map[int64]*entity.StudentShift{},
			onlyEnabled:   true,
			expect:        StudentShifts{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentShifts(tt.shifts, tt.studentShifts, tt.onlyEnabled))
		})
	}
}

func TestEnabledStudentShiftDetail(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name          string
		shifts        entity.Shifts
		date          time.Time
		studentShifts map[int64]*entity.StudentShift
		expect        *StudentShiftDetail
	}{
		{
			name: "success",
			shifts: entity.Shifts{
				{
					Shift: &lesson.Shift{
						Id:             1,
						ShiftSummaryId: 1,
						Date:           "20211226",
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				{
					Shift: &lesson.Shift{
						Id:             2,
						ShiftSummaryId: 1,
						Date:           "20211226",
						StartTime:      "1830",
						EndTime:        "2000",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			date: jst.Date(2021, 12, 26, 0, 0, 0, 0),
			studentShifts: map[int64]*entity.StudentShift{
				1: {
					StudentShift: &lesson.StudentShift{
						StudentId:      "studentid",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: &StudentShiftDetail{
				Date:     "20211226",
				IsClosed: false,
				Lessons: StudentShifts{
					{ID: 1, Enabled: true, StartTime: "1700", EndTime: "1830"},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewEnabledStudentShiftDetail(tt.shifts, tt.date, tt.studentShifts)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestEnabledStudentShiftDetails(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name          string
		summary       *ShiftSummary
		shifts        map[time.Time]entity.Shifts
		studentShifts map[int64]*entity.StudentShift
		expect        StudentShiftDetails
	}{
		{
			name: "success",
			summary: &ShiftSummary{
				ID:        1,
				Year:      2022,
				Month:     2,
				Status:    ShiftStatusAccepting,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			shifts: map[time.Time]entity.Shifts{
				jst.Date(2022, 2, 1, 0, 0, 0, 0): {
					{
						Shift: &lesson.Shift{
							Id:             1,
							ShiftSummaryId: 1,
							Date:           "20220201",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
					{
						Shift: &lesson.Shift{
							Id:             2,
							ShiftSummaryId: 1,
							Date:           "20220201",
							StartTime:      "1830",
							EndTime:        "2000",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
				jst.Date(2022, 2, 3, 0, 0, 0, 0): {
					{
						Shift: &lesson.Shift{
							Id:             3,
							ShiftSummaryId: 1,
							Date:           "20220203",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
				jst.Date(2022, 2, 4, 0, 0, 0, 0): {
					{
						Shift: &lesson.Shift{
							Id:             4,
							ShiftSummaryId: 1,
							Date:           "2022020",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
			},
			studentShifts: map[int64]*entity.StudentShift{
				1: {
					StudentShift: &lesson.StudentShift{
						StudentId:      "studentid",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				3: {
					StudentShift: &lesson.StudentShift{
						StudentId:      "studentid",
						ShiftId:        3,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: StudentShiftDetails{
				{
					Date:     "20220201",
					IsClosed: false,
					Lessons: StudentShifts{
						{ID: 1, Enabled: true, StartTime: "1700", EndTime: "1830"},
					},
				},
				{
					Date:     "20220203",
					IsClosed: false,
					Lessons: StudentShifts{
						{ID: 3, Enabled: true, StartTime: "1700", EndTime: "1830"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewEnabledStudentShiftDetails(tt.summary, tt.shifts, tt.studentShifts)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
