package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubmission(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		studentiD        string
		summaryID        int64
		decided          bool
		suggestedClasses int64
		expect           *StudentSubmission
	}{
		{
			name:             "success",
			studentiD:        "studentid",
			summaryID:        1,
			decided:          true,
			suggestedClasses: 8,
			expect: &StudentSubmission{
				StudentID:        "studentid",
				ShiftSummaryID:   1,
				Decided:          true,
				SuggestedClasses: 8,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewStudentSubmission(tt.studentiD, tt.summaryID, tt.decided, tt.suggestedClasses)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudentSubmission_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name       string
		submission *StudentSubmission
		expect     *lesson.StudentSubmission
	}{
		{
			name: "success",
			submission: &StudentSubmission{
				StudentID:        "studentid",
				ShiftSummaryID:   1,
				Decided:          true,
				SuggestedClasses: 8,
				CreatedAt:        now,
				UpdatedAt:        now,
			},
			expect: &lesson.StudentSubmission{
				StudentId:        "studentid",
				ShiftSummaryId:   1,
				Decided:          true,
				SuggestedClasses: 8,
				CreatedAt:        now.Unix(),
				UpdatedAt:        now.Unix(),
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

func TestStudentSubmissions_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions StudentSubmissions
		expect      []*lesson.StudentSubmission
	}{
		{
			name: "success",
			submissions: StudentSubmissions{
				{
					StudentID:        "studentid",
					ShiftSummaryID:   1,
					Decided:          true,
					SuggestedClasses: 8,
					CreatedAt:        now,
					UpdatedAt:        now,
				},
			},
			expect: []*lesson.StudentSubmission{
				{
					StudentId:        "studentid",
					ShiftSummaryId:   1,
					Decided:          true,
					SuggestedClasses: 8,
					CreatedAt:        now.Unix(),
					UpdatedAt:        now.Unix(),
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
