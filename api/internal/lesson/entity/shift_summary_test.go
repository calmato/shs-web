package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestShiftSummary(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		yearMonth int32
		openAt    int64
		endAt     int64
		expect    *ShiftSummary
	}{
		{
			name:      "success",
			yearMonth: 202202,
			openAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
			endAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
			expect: &ShiftSummary{
				YearMonth: 202202,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewShiftSummary(tt.yearMonth, tt.openAt, tt.endAt)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestShiftSummary_Year(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		summary *ShiftSummary
		expect  int
		isErr   bool
	}{
		{
			name: "success",
			summary: &ShiftSummary{
				YearMonth: 202202,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
			},
			expect: 2022,
			isErr:  false,
		},
		{
			name:    "failed to parse",
			summary: &ShiftSummary{YearMonth: -1},
			expect:  0,
			isErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.summary.Year()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestShiftSummary_Month(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		summary *ShiftSummary
		expect  int
		isErr   bool
	}{
		{
			name: "success",
			summary: &ShiftSummary{
				YearMonth: 202202,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
			},
			expect: 2,
			isErr:  false,
		},
		{
			name:    "failed to parse",
			summary: &ShiftSummary{YearMonth: -1},
			expect:  0,
			isErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual, err := tt.summary.Month()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestShiftSummary_Fill(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 02, 18, 30, 0, 0)
	tests := []struct {
		name         string
		now          time.Time
		summary      *ShiftSummary
		expectStatus ShiftStatus
	}{
		{
			name: "success to status waiting",
			now:  now,
			summary: &ShiftSummary{
				ID:        1,
				YearMonth: 202202,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectStatus: ShiftStatusWaiting,
		},
		{
			name: "success to status accepting",
			now:  now,
			summary: &ShiftSummary{
				ID:        1,
				YearMonth: 202201,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectStatus: ShiftStatusAccepting,
		},
		{
			name: "success to status finished",
			now:  now,
			summary: &ShiftSummary{
				ID:        1,
				YearMonth: 202112,
				Status:    ShiftStatusUnknown,
				OpenAt:    jst.Date(2021, 11, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2021, 11, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectStatus: ShiftStatusFinished,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.summary.Fill(tt.now)
			assert.Equal(t, tt.expectStatus, tt.summary.Status)
		})
	}
}

func TestShiftSummary_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 02, 18, 30, 0, 0)
	tests := []struct {
		name    string
		now     time.Time
		summary *ShiftSummary
		expect  *lesson.ShiftSummary
	}{
		{
			name: "success",
			now:  now,
			summary: &ShiftSummary{
				ID:        1,
				YearMonth: 202201,
				Status:    ShiftStatusAccepting,
				OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0),
				EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0),
				CreatedAt: now,
				UpdatedAt: now,
			},
			expect: &lesson.ShiftSummary{
				Id:        1,
				YearMonth: 202201,
				Status:    lesson.ShiftStatus_SHIFT_STATUS_ACCEPTING,
				OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0).Unix(),
				EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0).Unix(),
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.summary.Proto())
		})
	}
}

func TestShiftSummaries_Fill(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 02, 18, 30, 0, 0)
	tests := []struct {
		name      string
		now       time.Time
		summaries ShiftSummaries
		expect    ShiftSummaries
	}{
		{
			name: "success",
			now:  now,
			summaries: ShiftSummaries{
				{
					ID:        1,
					YearMonth: 202201,
					Status:    ShiftStatusUnknown,
					OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0),
					EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: ShiftSummaries{
				{
					ID:        1,
					YearMonth: 202201,
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
			tt.summaries.Fill(tt.now)
			assert.Equal(t, tt.expect, tt.summaries)
		})
	}
}

func TestShiftSummaries_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 12, 02, 18, 30, 0, 0)
	tests := []struct {
		name      string
		now       time.Time
		summaries ShiftSummaries
		expect    []*lesson.ShiftSummary
	}{
		{
			name: "success",
			now:  now,
			summaries: ShiftSummaries{
				{
					ID:        1,
					YearMonth: 202201,
					Status:    ShiftStatusAccepting,
					OpenAt:    jst.Date(2021, 12, 1, 0, 0, 0, 0),
					EndAt:     jst.Date(2021, 12, 15, 0, 0, 0, 0),
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []*lesson.ShiftSummary{
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.summaries.Proto())
		})
	}
}
