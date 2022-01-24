package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestStudentSubmission(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		studentiD string
		summaryID int64
		decided   bool
		lessons   SuggestedLessons
		expect    *StudentSubmission
	}{
		{
			name:      "success",
			studentiD: "studentid",
			summaryID: 1,
			decided:   true,
			lessons: SuggestedLessons{
				{SubjectID: 1, Total: 4},
				{SubjectID: 2, Total: 4},
			},
			expect: &StudentSubmission{
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
					{SubjectID: 2, Total: 4},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewStudentSubmission(tt.studentiD, tt.summaryID, tt.decided, tt.lessons)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudentSubmission_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		submission *StudentSubmission
		expect     *StudentSubmission
		isErr      bool
	}{
		{
			name: "success",
			submission: &StudentSubmission{
				StudentID:            "studentid",
				ShiftSummaryID:       1,
				Decided:              true,
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4},{"subjectId":2,"total":4}]`)),
			},
			expect: &StudentSubmission{
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
					{SubjectID: 2, Total: 4},
				},
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4},{"subjectId":2,"total":4}]`)),
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.submission.Fill()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.submission)
		})
	}
}

func TestStudentSubmission_FillJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		submission *StudentSubmission
		expect     *StudentSubmission
		isErr      bool
	}{
		{
			name: "success",
			submission: &StudentSubmission{
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
					{SubjectID: 2, Total: 4},
				},
			},
			expect: &StudentSubmission{
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
					{SubjectID: 2, Total: 4},
				},
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4},{"subjectId":2,"total":4}]`)),
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.submission.FillJSON()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.submission)
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
				StudentID:      "studentid",
				ShiftSummaryID: 1,
				Decided:        true,
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
					{SubjectID: 2, Total: 4},
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			expect: &lesson.StudentSubmission{
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
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, tt.submission.Proto())
		})
	}
}

func TestStudentSubmissions_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		submissions StudentSubmissions
		expect      StudentSubmissions
		isErr       bool
	}{
		{
			name: "success",
			submissions: StudentSubmissions{
				{
					StudentID:            "studentid",
					ShiftSummaryID:       1,
					Decided:              true,
					SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4},{"subjectId":2,"total":4}]`)),
				},
			},
			expect: StudentSubmissions{
				{
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: SuggestedLessons{
						{SubjectID: 1, Total: 4},
						{SubjectID: 2, Total: 4},
					},
					SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4},{"subjectId":2,"total":4}]`)),
				},
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.submissions.Fill()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.submissions)
		})
	}
}

func TestStudentSubmissions_FillJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		submissions StudentSubmissions
		expect      StudentSubmissions
		isErr       bool
	}{
		{
			name: "success",
			submissions: StudentSubmissions{
				{
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: SuggestedLessons{
						{SubjectID: 1, Total: 4},
						{SubjectID: 2, Total: 4},
					},
				},
			},
			expect: StudentSubmissions{
				{
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: SuggestedLessons{
						{SubjectID: 1, Total: 4},
						{SubjectID: 2, Total: 4},
					},
					SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4},{"subjectId":2,"total":4}]`)),
				},
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.submissions.FillJSON()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.submissions)
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
					StudentID:      "studentid",
					ShiftSummaryID: 1,
					Decided:        true,
					SuggestedLessons: SuggestedLessons{
						{SubjectID: 1, Total: 4},
						{SubjectID: 2, Total: 4},
					},
					CreatedAt: now,
					UpdatedAt: now,
				},
			},
			expect: []*lesson.StudentSubmission{
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

func TestSuggestedLesson(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		lesson *lesson.StudentSuggestedLesson
		expect *SuggestedLesson
	}{
		{
			name: "success",
			lesson: &lesson.StudentSuggestedLesson{
				SubjectId: 1,
				Total:     4,
			},
			expect: &SuggestedLesson{
				SubjectID: 1,
				Total:     4,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSuggestedLesson(tt.lesson))
		})
	}
}

func TestSuggestedLessons(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		lessons []*lesson.StudentSuggestedLesson
		expect  SuggestedLessons
	}{
		{
			name: "success",
			lessons: []*lesson.StudentSuggestedLesson{
				{
					SubjectId: 1,
					Total:     4,
				},
			},
			expect: SuggestedLessons{
				{
					SubjectID: 1,
					Total:     4,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewSuggestedLessons(tt.lessons))
		})
	}
}
