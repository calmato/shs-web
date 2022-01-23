package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestTeacherShift(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shift  *lesson.TeacherShift
		expect *TeacherShift
	}{
		{
			name: "success",
			shift: &lesson.TeacherShift{
				TeacherId:      "teacherid",
				ShiftId:        1,
				ShiftSummaryId: 1,
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
			expect: &TeacherShift{
				TeacherShift: &lesson.TeacherShift{
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
			assert.Equal(t, tt.expect, NewTeacherShift(tt.shift))
		})
	}
}

func TestTeacherShifts(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts []*lesson.TeacherShift
		expect TeacherShifts
	}{
		{
			name: "success",
			shifts: []*lesson.TeacherShift{
				{
					TeacherId:      "teacherid",
					ShiftId:        1,
					ShiftSummaryId: 1,
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: TeacherShifts{
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "teacherid",
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
			assert.Equal(t, tt.expect, NewTeacherShifts(tt.shifts))
		})
	}
}

func TestTeacherShifts_TeacherIDs(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts TeacherShifts
		expect []string
	}{
		{
			name: "success",
			shifts: TeacherShifts{
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "teacherid1",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "teacherid1",
						ShiftId:        2,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "teacherid2",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: []string{
				"teacherid1",
				"teacherid2",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.shifts.TeacherIDs()
			assert.Len(t, actual, len(tt.expect))
			for i := range tt.expect {
				assert.Contains(t, actual, tt.expect[i])
			}
		})
	}
}

func TestTeacherShifts_MapByShiftID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shifts TeacherShifts
		expect map[int64]*TeacherShift
	}{
		{
			name: "success",
			shifts: TeacherShifts{
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "teacherid",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: map[int64]*TeacherShift{
				1: {
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "teacherid",
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
