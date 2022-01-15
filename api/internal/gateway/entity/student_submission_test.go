package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubmission(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name       string
		submission *lesson.StudentSubmission
		expect     *StudentSubmission
	}{
		{
			name: "success",
			submission: &lesson.StudentSubmission{
				StudentId:        "studentid",
				ShiftSummaryId:   1,
				Decided:          true,
				SuggestedClasses: 8,
				CreatedAt:        now.Unix(),
				UpdatedAt:        now.Unix(),
			},
			expect: &StudentSubmission{
				StudentSubmission: &lesson.StudentSubmission{
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
			assert.Equal(t, tt.expect, NewStudentSubmission(tt.submission))
		})
	}
}

func TestStudentSubmissions(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions []*lesson.StudentSubmission
		expect      StudentSubmissions
	}{
		{
			name: "success",
			submissions: []*lesson.StudentSubmission{
				{
					StudentId:        "studentid",
					ShiftSummaryId:   1,
					Decided:          true,
					SuggestedClasses: 8,
					CreatedAt:        now.Unix(),
					UpdatedAt:        now.Unix(),
				},
			},
			expect: StudentSubmissions{
				{
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid",
						ShiftSummaryId:   1,
						Decided:          true,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubmissions(tt.submissions))
		})
	}
}

func TestStudentSubmissions_MapByShiftSummaryID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions StudentSubmissions
		expect      map[int64]*StudentSubmission
	}{
		{
			name: "success",
			submissions: StudentSubmissions{
				{
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid",
						ShiftSummaryId:   1,
						Decided:          true,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
				{
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid",
						ShiftSummaryId:   2,
						Decided:          false,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
			},
			expect: map[int64]*StudentSubmission{
				1: {
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid",
						ShiftSummaryId:   1,
						Decided:          true,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
				2: {
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid",
						ShiftSummaryId:   2,
						Decided:          false,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.submissions.MapByShiftSummaryID())
		})
	}
}

func TestStudentSubmissions_MapByStudentID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions StudentSubmissions
		expect      map[string]*StudentSubmission
	}{
		{
			name: "success",
			submissions: StudentSubmissions{
				{
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid1",
						ShiftSummaryId:   1,
						Decided:          true,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
				{
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid2",
						ShiftSummaryId:   1,
						Decided:          false,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
			},
			expect: map[string]*StudentSubmission{
				"studentid1": {
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid1",
						ShiftSummaryId:   1,
						Decided:          true,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
				"studentid2": {
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "studentid2",
						ShiftSummaryId:   1,
						Decided:          false,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.submissions.MapByStudentID())
		})
	}
}
