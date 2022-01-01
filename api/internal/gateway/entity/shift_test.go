package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name   string
		shift  *lesson.Shift
		expect *Shift
	}{
		{
			name: "success",
			shift: &lesson.Shift{
				Id:             1,
				ShiftSummaryId: 1,
				Date:           "20211226",
				StartTime:      "1700",
				EndTime:        "1830",
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
			expect: &Shift{
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
	now := jst.Now()
	tests := []struct {
		name   string
		shifts []*lesson.Shift
		expect Shifts
	}{
		{
			name: "success",
			shifts: []*lesson.Shift{
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
			expect: Shifts{
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
	now := jst.Now()
	tests := []struct {
		name   string
		shifts Shifts
		expect map[time.Time]Shifts
		isErr  bool
	}{
		{
			name: "success",
			shifts: Shifts{
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
				{
					Shift: &lesson.Shift{
						Id:             3,
						ShiftSummaryId: 1,
						Date:           "20211227",
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: map[time.Time]Shifts{
				jst.Date(2021, 12, 26, 0, 0, 0, 0): {
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
				jst.Date(2021, 12, 27, 0, 0, 0, 0): {
					{
						Shift: &lesson.Shift{
							Id:             3,
							ShiftSummaryId: 1,
							Date:           "20211227",
							StartTime:      "1700",
							EndTime:        "1830",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
			},
			isErr: false,
		},
		{
			name: "failed to parse date",
			shifts: Shifts{
				{
					Shift: &lesson.Shift{
						Id:             1,
						ShiftSummaryId: 1,
						Date:           "20211200",
						StartTime:      "1700",
						EndTime:        "1830",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: nil,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.shifts.GroupByDate()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
