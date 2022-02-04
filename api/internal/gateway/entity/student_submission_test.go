package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubmission(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 2, 1, 18, 30, 0, 0)
	tests := []struct {
		name       string
		submission *lesson.StudentSubmission
		expect     *StudentSubmission
	}{
		{
			name: "success",
			submission: &lesson.StudentSubmission{
				StudentId:      "studentid",
				ShiftSummaryId: 1,
				Decided:        true,
				SuggestedLessons: []*lesson.SuggestedLesson{
					{SubjectId: 1, Total: 4},
					{SubjectId: 2, Total: 4},
				},
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
			expect: &StudentSubmission{
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				SuggestedLessons: StudentSuggestedLessons{
					{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
					{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
				},
				CreatedAt: now,
				UpdatedAt: now,
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
	now := jst.Date(2022, 1, 2, 18, 30, 0, 0)
	tests := []struct {
		name        string
		submissions []*lesson.StudentSubmission
		expect      StudentSubmissions
	}{
		{
			name: "success",
			submissions: []*lesson.StudentSubmission{
				{
					StudentId:      "studentid",
					ShiftSummaryId: 1,
					Decided:        true,
					SuggestedLessons: []*lesson.SuggestedLesson{
						{SubjectId: 1, Total: 4},
						{SubjectId: 2, Total: 4},
					},
					CreatedAt: now.Unix(),
					UpdatedAt: now.Unix(),
				},
			},
			expect: StudentSubmissions{
				{
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
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
			assert.Equal(t, tt.expect, NewStudentSubmissions(tt.submissions))
		})
	}
}

func TestStudentSubmissions_MapByShiftSummaryID(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 2, 18, 30, 0, 0)
	tests := []struct {
		name        string
		submissions StudentSubmissions
		expect      map[int64]*StudentSubmission
	}{
		{
			name: "success",
			submissions: StudentSubmissions{
				{
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID:      "studentid",
					ShiftSummaryID: 2,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: map[int64]*StudentSubmission{
				1: {
					StudentID:      "studentid",
					ShiftSummaryID: 2,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				2: {
					StudentID:      "studentid",
					ShiftSummaryID: 2,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
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
					StudentID:      "studentid1",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					StudentID:      "studentid2",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: map[string]*StudentSubmission{
				"studentid1": {
					StudentID:      "studentid1",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
				"studentid2": {
					StudentID:      "studentid2",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: StudentSuggestedLessons{
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
						{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
					},
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
			assert.Equal(t, tt.expect, tt.submissions.MapByStudentID())
		})
	}
}
