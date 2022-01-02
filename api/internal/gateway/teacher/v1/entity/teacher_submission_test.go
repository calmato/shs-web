package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestTeacherSubmission(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name       string
		summary    *entity.ShiftSummary
		submission *entity.TeacherSubmission
		expect     *TeacherSubmission
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
			submission: &entity.TeacherSubmission{
				TeacherSubmission: &lesson.TeacherSubmission{
					TeacherId:      "teacherid",
					ShiftSummaryId: 1,
					Decided:        true,
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: &TeacherSubmission{
				ID:               1,
				Year:             2022,
				Month:            1,
				ShiftStatus:      ShiftStatusAccepting,
				SubmissionStatus: TeacherSubmissionStatusSubmitted,
				OpenAt:           jst.Date(2021, 12, 1, 0, 0, 0, 0),
				EndAt:            jst.Date(2021, 12, 15, 0, 0, 0, 0),
				CreatedAt:        now,
				UpdatedAt:        now,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmission(tt.summary, tt.submission))
		})
	}
}

func TestTeacherSubmissions(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name        string
		summaries   entity.ShiftSummaries
		submissions map[int64]*entity.TeacherSubmission
		expect      TeacherSubmissions
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
				{
					ShiftSummary: &lesson.ShiftSummary{
						Id:        2,
						YearMonth: 202202,
						Status:    lesson.ShiftStatus_SHIFT_STATUS_WAITING,
						OpenAt:    jst.Date(2022, 1, 1, 0, 0, 0, 0).Unix(),
						EndAt:     jst.Date(2022, 1, 15, 0, 0, 0, 0).Unix(),
						CreatedAt: now.Unix(),
						UpdatedAt: now.Unix(),
					},
				},
			},
			submissions: map[int64]*entity.TeacherSubmission{
				1: {
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: TeacherSubmissions{
				{
					ID:               1,
					Year:             2022,
					Month:            1,
					ShiftStatus:      ShiftStatusAccepting,
					SubmissionStatus: TeacherSubmissionStatusSubmitted,
					OpenAt:           jst.Date(2021, 12, 1, 0, 0, 0, 0),
					EndAt:            jst.Date(2021, 12, 15, 0, 0, 0, 0),
					CreatedAt:        now,
					UpdatedAt:        now,
				},
				{
					ID:               2,
					Year:             2022,
					Month:            2,
					ShiftStatus:      ShiftStatusWaiting,
					SubmissionStatus: TeacherSubmissionStatusWaiting,
					OpenAt:           jst.Date(2022, 1, 1, 0, 0, 0, 0),
					EndAt:            jst.Date(2022, 1, 15, 0, 0, 0, 0),
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissions(tt.summaries, tt.submissions))
		})
	}
}

func TestTeacherSubmissionStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		summary    *entity.ShiftSummary
		submission *entity.TeacherSubmission
		expect     TeacherSubmissionStatus
	}{
		{
			name: "waiting status when not decided",
			summary: &entity.ShiftSummary{
				ShiftSummary: &lesson.ShiftSummary{Id: 1},
			},
			submission: &entity.TeacherSubmission{
				TeacherSubmission: &lesson.TeacherSubmission{Decided: false},
			},
			expect: TeacherSubmissionStatusWaiting,
		},
		{
			name: "waiting status when submission is nil",
			summary: &entity.ShiftSummary{
				ShiftSummary: &lesson.ShiftSummary{Id: 1},
			},
			submission: nil,
			expect:     TeacherSubmissionStatusWaiting,
		},
		{
			name: "submitted status",
			summary: &entity.ShiftSummary{
				ShiftSummary: &lesson.ShiftSummary{Id: 1},
			},
			submission: &entity.TeacherSubmission{
				TeacherSubmission: &lesson.TeacherSubmission{Decided: true},
			},
			expect: TeacherSubmissionStatusSubmitted,
		},
		{
			name:       "unknown status",
			summary:    nil,
			submission: nil,
			expect:     TeacherSubmissionStatusUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissionStatus(tt.summary, tt.submission))
		})
	}
}
