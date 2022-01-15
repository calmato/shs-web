package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentShift(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shift  *lesson.StudentShift
		expect *StudentShift
	}{
		{
			name: "success",
			shift: &lesson.StudentShift{
				StudentId:      "studentid",
				ShiftId:        1,
				ShiftSummaryId: 1,
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
			expect: &StudentShift{
				StudentShift: &lesson.StudentShift{
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
			assert.Equal(t, tt.expect, NewStudentShift(tt.shift))
		})
	}
}

func TestStudentShifts(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts []*lesson.StudentShift
		expect StudentShifts
	}{
		{
			name: "success",
			shifts: []*lesson.StudentShift{
				{
					StudentId:      "studentid",
					ShiftId:        1,
					ShiftSummaryId: 1,
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: StudentShifts{
				{
					StudentShift: &lesson.StudentShift{
						StudentId:      "studentid",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentShifts(tt.shifts))
		})
	}
}

func TestStudentShifts_MapByShiftID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts StudentShifts
		expect map[int64]*StudentShift
	}{
		{
			name: "success",
			shifts: StudentShifts{
				{
					StudentShift: &lesson.StudentShift{
						StudentId:      "studentid",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: map[int64]*StudentShift{
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.shifts.MapByShiftID())
		})
	}
}
