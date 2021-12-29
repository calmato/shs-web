package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShiftSummary(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name    string
		summary *entity.ShiftSummary
		expect  *ShiftSummary
	}{
		{
			name: "success",
			summary: &entity.ShiftSummary{
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
			expect: &ShiftSummary{
				ID:        1,
				Year:      2022,
				Month:     1,
				Status:    ShiftStatusAccepting,
				OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
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
	now := jst.Date(2021, 12, 1, 12, 30, 0, 0)
	tests := []struct {
		name      string
		summaries entity.ShiftSummaries
		expect    ShiftSummaries
	}{
		{
			name: "success",
			summaries: entity.ShiftSummaries{
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
			expect: ShiftSummaries{
				{
					ID:        1,
					Year:      2022,
					Month:     1,
					Status:    ShiftStatusAccepting,
					OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0),
					EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
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

func TestShiftStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		status lesson.ShiftStatus
		expect ShiftStatus
	}{
		{
			name:   "success waiting status",
			status: lesson.ShiftStatus_SHIFT_STATUS_WAITING,
			expect: ShiftStatusWaiting,
		},
		{
			name:   "success accepting status",
			status: lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
			expect: ShiftStatusAccepting,
		},
		{
			name:   "success finished status",
			status: lesson.ShiftStatus_SHIFT_STATUS_FINISHED,
			expect: ShiftStatusFinished,
		},
		{
			name:   "success invalid status",
			status: -1,
			expect: ShiftStatusUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewShiftStatus(tt.status))
		})
	}
}

func TestShiftStatus_LessonShiftStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		status ShiftStatus
		expect lesson.ShiftStatus
		isErr  bool
	}{
		{
			name:   "success waiting status",
			status: ShiftStatusWaiting,
			expect: lesson.ShiftStatus_SHIFT_STATUS_WAITING,
			isErr:  false,
		},
		{
			name:   "success accepting status",
			status: ShiftStatusAccepting,
			expect: lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
			isErr:  false,
		},
		{
			name:   "success finished status",
			status: ShiftStatusFinished,
			expect: lesson.ShiftStatus_SHIFT_STATUS_FINISHED,
			isErr:  false,
		},
		{
			name:   "failed to invalid shift status",
			status: ShiftStatusUnknown,
			expect: lesson.ShiftStatus_SHIFT_STATUS_UNKNOWN,
			isErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.status.LessonShiftStatus()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}
