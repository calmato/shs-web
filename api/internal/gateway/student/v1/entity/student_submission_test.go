package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubmission(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name       string
		summary    *entity.ShiftSummary
		submission *entity.StudentSubmission
		expect     *StudentSubmission
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
			submission: &entity.StudentSubmission{
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expect: &StudentSubmission{
				ShiftSummaryID:   1,
				Year:             2022,
				Month:            1,
				ShiftStatus:      ShiftStatusAccepting,
				SubmissionStatus: StudentSubmissionStatusSubmitted,
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
			assert.Equal(t, tt.expect, NewStudentSubmission(tt.summary, tt.submission))
		})
	}
}

func TestStudentSubmissions(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 1, 0, 0, 0, 0)
	tests := []struct {
		name        string
		summaries   entity.ShiftSummaries
		submissions map[int64]*entity.StudentSubmission
		expect      StudentSubmissions
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
			submissions: map[int64]*entity.StudentSubmission{
				1: {
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: StudentSubmissions{
				{
					ShiftSummaryID:   1,
					Year:             2022,
					Month:            1,
					ShiftStatus:      ShiftStatusAccepting,
					SubmissionStatus: StudentSubmissionStatusSubmitted,
					OpenAt:           jst.Date(2021, 12, 1, 0, 0, 0, 0),
					EndAt:            jst.Date(2021, 12, 15, 0, 0, 0, 0),
					CreatedAt:        now,
					UpdatedAt:        now,
				},
				{
					ShiftSummaryID:   2,
					Year:             2022,
					Month:            2,
					ShiftStatus:      ShiftStatusWaiting,
					SubmissionStatus: StudentSubmissionStatusWaiting,
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
			assert.Equal(t, tt.expect, NewStudentSubmissions(tt.summaries, tt.submissions))
		})
	}
}

func TestStudentSubmissionStatus(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		summary    *entity.ShiftSummary
		submission *entity.StudentSubmission
		expect     StudentSubmissionStatus
	}{
		{
			name:       "student submission status waiting",
			summary:    &entity.ShiftSummary{},
			submission: nil,
			expect:     StudentSubmissionStatusWaiting,
		},
		{
			name:       "student submission status submitted",
			summary:    &entity.ShiftSummary{},
			submission: &entity.StudentSubmission{Decided: true},
			expect:     StudentSubmissionStatusSubmitted,
		},
		{
			name:       "student submission status unknown",
			summary:    nil,
			submission: nil,
			expect:     StudentSubmissionStatusUnknown,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubmissionStatus(tt.summary, tt.submission))
		})
	}
}
