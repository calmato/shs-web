package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShiftSummary(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name    string
		summary *lesson.ShiftSummary
		expect  *ShiftSummary
	}{
		{
			name: "success",
			summary: &lesson.ShiftSummary{
				Id:        1,
				YearMonth: 202201,
				Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
				OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0).Unix(),
				EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0).Unix(),
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
			expect: &ShiftSummary{
				ShiftSummary: &lesson.ShiftSummary{
					Id:        1,
					YearMonth: 202201,
					Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0).Unix(),
					EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0).Unix(),
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
			assert.Equal(t, tt.expect, NewShiftSummary(tt.summary))
		})
	}
}

func TestShiftSummaries(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name      string
		summaries []*lesson.ShiftSummary
		expect    ShiftSummaries
	}{
		{
			name: "success",
			summaries: []*lesson.ShiftSummary{
				{
					Id:        1,
					YearMonth: 202201,
					Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
					OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0).Unix(),
					EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0).Unix(),
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				},
			},
			expect: ShiftSummaries{
				{
					ShiftSummary: &lesson.ShiftSummary{
						Id:        1,
						YearMonth: 202201,
						Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
						OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0).Unix(),
						EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0).Unix(),
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewShiftSummaries(tt.summaries))
		})
	}
}
