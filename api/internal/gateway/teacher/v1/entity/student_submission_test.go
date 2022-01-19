package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
	"github.com/stretchr/testify/assert"
)

func TestStudentSubmissionDetail(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name       string
		student    *entity.Student
		submission *entity.StudentSubmission
		lessons    entity.Lessons
		expect     *StudentSubmissionDetail
	}{
		{
			name: "success",
			student: &entity.Student{
				Student: &user.Student{
					Id:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "student-test001@calmato.jp",
					SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
					Grade:         3,
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
			submission: &entity.StudentSubmission{
				StudentSubmission: &lesson.StudentSubmission{
					StudentId:        "kSByoE6FetnPs5Byk3a9Zx",
					ShiftSummaryId:   1,
					Decided:          true,
					SuggestedClasses: 8,
					CreatedAt:        now.Unix(),
					UpdatedAt:        now.Unix(),
				},
			},
			lessons: entity.Lessons{
				{
					Lesson: &lesson.Lesson{
						Id:             1,
						ShiftSummaryId: 1,
						ShiftId:        1,
						SubjectId:      1,
						RoomId:         1,
						TeacherId:      "teacherid",
						StudentId:      "studentid",
						Notes:          "感想",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: &StudentSubmissionDetail{
				Student: &Student{
					ID:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "student-test001@calmato.jp",
					SchoolType:    SchoolTypeHighSchool,
					Grade:         3,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
				LessonTotal:           1,
				SuggestedClassesTotal: 8,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubmissionDetail(tt.student, tt.submission, tt.lessons))
		})
	}
}

func TestStudentSubmissionDetails(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name        string
		students    entity.Students
		submissions map[string]*entity.StudentSubmission
		lessons     map[string]entity.Lessons
		expect      StudentSubmissionDetails
	}{
		{
			name: "success",
			students: entity.Students{
				{
					Student: &user.Student{
						Id:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "student-test001@calmato.jp",
						SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						Grade:         3,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
				{
					Student: &user.Student{
						Id:            "studentid",
						LastName:      "テスト",
						FirstName:     "講師",
						LastNameKana:  "てすと",
						FirstNameKana: "こうし",
						Mail:          "student-test002@calmato.jp",
						SchoolType:    user.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						Grade:         3,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
			submissions: map[string]*entity.StudentSubmission{
				"kSByoE6FetnPs5Byk3a9Zx": {
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:        "kSByoE6FetnPs5Byk3a9Zx",
						ShiftSummaryId:   1,
						Decided:          true,
						SuggestedClasses: 8,
						CreatedAt:        now.Unix(),
						UpdatedAt:        now.Unix(),
					},
				},
			},
			lessons: map[string]entity.Lessons{
				"studentid": {
					{
						Lesson: &lesson.Lesson{
							Id:             1,
							ShiftSummaryId: 1,
							ShiftId:        1,
							SubjectId:      1,
							RoomId:         1,
							TeacherId:      "teacherid",
							StudentId:      "studentid",
							Notes:          "感想",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
			},
			expect: StudentSubmissionDetails{
				{
					Student: &Student{
						ID:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "student-test001@calmato.jp",
						SchoolType:    SchoolTypeHighSchool,
						Grade:         3,
						CreatedAt:     now,
						UpdatedAt:     now,
					},
					LessonTotal:           0,
					SuggestedClassesTotal: 8,
				},
				{
					Student: &Student{
						ID:            "studentid",
						LastName:      "テスト",
						FirstName:     "講師",
						LastNameKana:  "てすと",
						FirstNameKana: "こうし",
						Mail:          "student-test002@calmato.jp",
						SchoolType:    SchoolTypeHighSchool,
						Grade:         3,
						CreatedAt:     now,
						UpdatedAt:     now,
					},
					LessonTotal:           1,
					SuggestedClassesTotal: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSubmissionDetails(tt.students, tt.submissions, tt.lessons))
		})
	}
}
