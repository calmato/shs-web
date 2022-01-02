package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestTeacherSubmission(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		teacheriD string
		summaryID int64
		decided   bool
		expect    *TeacherSubmission
	}{
		{
			name:      "success",
			teacheriD: "teacherid",
			summaryID: 1,
			decided:   true,
			expect: &TeacherSubmission{
				TeacherID:      "teacherid",
				ShiftSummaryID: 1,
				Decided:        true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewTeacherSubmission(tt.teacheriD, tt.summaryID, tt.decided)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestTeacherSubmission_Fill(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name       string
		submission *TeacherSubmission
		shifts     TeacherShifts
		expect     TeacherSubmissionStatus
	}{
		{
			name: "success to status submitted",
			submission: &TeacherSubmission{
				TeacherID:      "teacherid",
				ShiftSummaryID: 1,
				Decided:        true,
				Status:         TeacherSubmissionStatusUnknown,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			shifts: TeacherShifts{},
			expect: TeacherSubmissionStatusSubmitted,
		},
		{
			name: "success to status editing",
			submission: &TeacherSubmission{
				TeacherID:      "teacherid",
				ShiftSummaryID: 1,
				Decided:        false,
				Status:         TeacherSubmissionStatusUnknown,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			shifts: TeacherShifts{
				{
					TeacherID:      "teacherid",
					ShiftSummaryID: 1,
					ShiftID:        1,
				},
			},
			expect: TeacherSubmissionStatusEditing,
		},
		{
			name: "success to status waiting",
			submission: &TeacherSubmission{
				TeacherID:      "teacherid",
				ShiftSummaryID: 1,
				Decided:        false,
				Status:         TeacherSubmissionStatusUnknown,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			shifts: TeacherShifts{},
			expect: TeacherSubmissionStatusWaiting,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.submission.Fill(tt.shifts)
			assert.Equal(t, tt.expect, tt.submission.Status)
		})
	}
}

func TestTeacherSubmission_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name       string
		submission *TeacherSubmission
		expect     *lesson.TeacherSubmission
	}{
		{
			name: "success",
			submission: &TeacherSubmission{
				TeacherID:      "teacherid",
				ShiftSummaryID: 1,
				Decided:        true,
				Status:         TeacherSubmissionStatusEditing,
				CreatedAt:      now,
				UpdatedAt:      now,
			},
			expect: &lesson.TeacherSubmission{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				Decided:        true,
				Status:         lesson.TeacherSubmissionStatus_TEACHER_SUBMISSION_STATUS_EDITING,
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.submission.Proto())
		})
	}
}

func TestTeacherSubmissions_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions TeacherSubmissions
		expect      []*lesson.TeacherSubmission
	}{
		{
			name: "success",
			submissions: TeacherSubmissions{
				{
					TeacherID:      "teacherid",
					ShiftSummaryID: 1,
					Decided:        true,
					Status:         TeacherSubmissionStatusEditing,
					CreatedAt:      now,
					UpdatedAt:      now,
				},
			},
			expect: []*lesson.TeacherSubmission{
				{
					TeacherId:      "teacherid",
					ShiftSummaryId: 1,
					Decided:        true,
					Status:         lesson.TeacherSubmissionStatus_TEACHER_SUBMISSION_STATUS_EDITING,
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
			assert.Equal(t, tt.expect, tt.submissions.Proto())
		})
	}
}
