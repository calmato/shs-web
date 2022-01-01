package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestTeacherShift(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		teacherID string
		summaryID int64
		shiftID   int64
		expect    *TeacherShift
	}{
		{
			name:      "success",
			teacherID: "teacherid",
			summaryID: 1,
			shiftID:   1,
			expect: &TeacherShift{
				TeacherID:      "teacherid",
				ShiftID:        1,
				ShiftSummaryID: 1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewTeacherShift(tt.teacherID, tt.summaryID, tt.shiftID)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestTeacherShifts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		teacherID string
		summaryID int64
		shiftIDs  []int64
		expect    TeacherShifts
	}{
		{
			name:      "success",
			teacherID: "teacherid",
			summaryID: 1,
			shiftIDs:  []int64{1, 2},
			expect: TeacherShifts{
				{
					TeacherID:      "teacherid",
					ShiftID:        1,
					ShiftSummaryID: 1,
				},
				{
					TeacherID:      "teacherid",
					ShiftID:        2,
					ShiftSummaryID: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewTeacherShifts(tt.teacherID, tt.summaryID, tt.shiftIDs)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestTeacherShift_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shift  *TeacherShift
		expect *lesson.TeacherShift
	}{
		{
			name: "success",
			shift: &TeacherShift{
				TeacherID:      "teacherid",
				ShiftID:        1,
				ShiftSummaryID: 1,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expect: &lesson.TeacherShift{
				TeacherId:      "teacherid",
				ShiftId:        1,
				ShiftSummaryId: 1,
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

func TestTeacherShifts_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts TeacherShifts
		expect []*lesson.TeacherShift
	}{
		{
			name: "success",
			shifts: TeacherShifts{
				{
					TeacherID:      "teacherid",
					ShiftID:        1,
					ShiftSummaryID: 1,
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []*lesson.TeacherShift{
				{
					TeacherId:      "teacherid",
					ShiftId:        1,
					ShiftSummaryId: 1,
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
