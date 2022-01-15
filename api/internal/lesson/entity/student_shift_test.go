package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentShift(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		studentID string
		summaryID int64
		shiftID   int64
		expect    *StudentShift
	}{
		{
			name:      "success",
			studentID: "studentid",
			summaryID: 1,
			shiftID:   1,
			expect: &StudentShift{
				StudentID:      "studentid",
				ShiftID:        1,
				ShiftSummaryID: 1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewStudentShift(tt.studentID, tt.summaryID, tt.shiftID)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudentShifts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		studentID string
		summaryID int64
		shiftIDs  []int64
		expect    StudentShifts
	}{
		{
			name:      "success",
			studentID: "studentid",
			summaryID: 1,
			shiftIDs:  []int64{1, 2},
			expect: StudentShifts{
				{
					StudentID:      "studentid",
					ShiftID:        1,
					ShiftSummaryID: 1,
				},
				{
					StudentID:      "studentid",
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
			actual := NewStudentShifts(tt.studentID, tt.summaryID, tt.shiftIDs)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudentShift_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shift  *StudentShift
		expect *lesson.StudentShift
	}{
		{
			name: "success",
			shift: &StudentShift{
				StudentID:      "studentid",
				ShiftID:        1,
				ShiftSummaryID: 1,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expect: &lesson.StudentShift{
				StudentId:      "studentid",
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

func TestStudentShifts_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts StudentShifts
		expect []*lesson.StudentShift
	}{
		{
			name: "success",
			shifts: StudentShifts{
				{
					StudentID:      "studentid",
					ShiftID:        1,
					ShiftSummaryID: 1,
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []*lesson.StudentShift{
				{
					StudentId:      "studentid",
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
