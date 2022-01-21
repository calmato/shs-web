package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/calmato/shs-web/api/proto/user"
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
				StudentSubmission: &lesson.StudentSubmission{
					StudentId:      "studentid",
					ShiftSummaryId: 1,
					Decided:        true,
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
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
					StudentSubmission: &lesson.StudentSubmission{
						StudentId:      "studentid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
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

func TestStudentSubmissionDetail(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name       string
		student    *entity.Student
		subjects   entity.Subjects
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
			subjects: entity.Subjects{
				{
					Subject: &classroom.Subject{
						Id:         1,
						Name:       "国語",
						Color:      "#F8BBD0",
						SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						CreatedAt:  now.Unix(),
						UpdatedAt:  now.Unix(),
					},
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
					Subjects: Subjects{
						{
							ID:         1,
							Name:       "国語",
							Color:      "#F8BBD0",
							SchoolType: SchoolTypeHighSchool,
							CreatedAt:  now,
							UpdatedAt:  now,
						},
					},
					CreatedAt: now,
					UpdatedAt: now,
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
			assert.Equal(t, tt.expect, NewStudentSubmissionDetail(tt.student, tt.subjects, tt.submission, tt.lessons))
		})
	}
}

func TestStudentSubmissionDetails(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name        string
		students    entity.Students
		subjects    map[string]entity.Subjects
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
			subjects: map[string]entity.Subjects{
				"kSByoE6FetnPs5Byk3a9Zx": {
					{
						Subject: &classroom.Subject{
							Id:         1,
							Name:       "国語",
							Color:      "#F8BBD0",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
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
						Subjects: Subjects{
							{
								ID:         1,
								Name:       "国語",
								Color:      "#F8BBD0",
								SchoolType: SchoolTypeHighSchool,
								CreatedAt:  now,
								UpdatedAt:  now,
							},
						},
						CreatedAt: now,
						UpdatedAt: now,
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
						Subjects:      Subjects{},
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
			assert.Equal(t, tt.expect, NewStudentSubmissionDetails(tt.students, tt.subjects, tt.submissions, tt.lessons))
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
			submission: &entity.StudentSubmission{StudentSubmission: &lesson.StudentSubmission{Decided: true}},
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
