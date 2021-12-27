package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

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
