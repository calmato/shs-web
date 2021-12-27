package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name   string
		shift  *entity.Shift
		expect *Shift
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
			expect: &Shift{
				ID:        1,
				Date:      "20211226",
				StartTime: "1700",
				EndTime:   "1830",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewShift(tt.shift))
		})
	}
}

func TestShifts(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name   string
		shifts entity.Shifts
		expect Shifts
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
			expect: Shifts{
				{
					ID:        1,
					Date:      "20211226",
					StartTime: "1700",
					EndTime:   "1830",
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewShifts(tt.shifts))
		})
	}
}

func TestShifts_GroupByDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		shifts Shifts
		expect map[string]Shifts
	}{
		{
			name: "success",
			shifts: Shifts{
				{
					ID:        1,
					Date:      "20211226",
					StartTime: "1700",
					EndTime:   "1830",
				},
				{
					ID:        2,
					Date:      "20211226",
					StartTime: "1830",
					EndTime:   "2000",
				},
				{
					ID:        3,
					Date:      "20211227",
					StartTime: "1700",
					EndTime:   "1830",
				},
			},
			expect: map[string]Shifts{
				"20211226": {
					{
						ID:        1,
						Date:      "20211226",
						StartTime: "1700",
						EndTime:   "1830",
					},
					{
						ID:        2,
						Date:      "20211226",
						StartTime: "1830",
						EndTime:   "2000",
					},
				},
				"20211227": {
					{
						ID:        3,
						Date:      "20211227",
						StartTime: "1700",
						EndTime:   "1830",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.shifts.GroupByDate())
		})
	}
}
