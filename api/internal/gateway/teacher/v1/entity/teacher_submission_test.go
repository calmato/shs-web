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
				ShiftSummaryID:   1,
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
					ShiftSummaryID:   1,
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
					ShiftSummaryID:   2,
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

func TestTeacherSubmissionDetail(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name       string
		teacher    *entity.Teacher
		subjects   entity.Subjects
		submission *entity.TeacherSubmission
		lessons    entity.Lessons
		expect     *TeacherSubmissionDetail
	}{
		{
			name: "success",
			teacher: &entity.Teacher{
				Teacher: &user.Teacher{
					Id:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "??????",
					FirstName:     "??????",
					LastNameKana:  "????????????",
					FirstNameKana: "????????????",
					Mail:          "teacher-test001@calmato.jp",
					Role:          user.Role_ROLE_TEACHER,
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
			subjects: entity.Subjects{
				{
					Subject: &classroom.Subject{
						Id:         1,
						Name:       "??????",
						Color:      "#F8BBD0",
						SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
						CreatedAt:  now.Unix(),
						UpdatedAt:  now.Unix(),
					},
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
						Notes:          "??????",
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: &TeacherSubmissionDetail{
				Teacher: &Teacher{
					ID:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "??????",
					FirstName:     "??????",
					LastNameKana:  "????????????",
					FirstNameKana: "????????????",
					Mail:          "teacher-test001@calmato.jp",
					Role:          RoleTeacher,
					CreatedAt:     now,
					UpdatedAt:     now,
					Subjects: map[SchoolType]Subjects{
						SchoolTypeElementarySchool: {},
						SchoolTypeJuniorHighSchool: {},
						SchoolTypeHighSchool: {
							{
								ID:         1,
								Name:       "??????",
								Color:      "#F8BBD0",
								SchoolType: SchoolTypeHighSchool,
								CreatedAt:  now,
								UpdatedAt:  now,
							},
						},
					},
				},
				IsSubmit:    true,
				LessonTotal: 1,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissionDetail(tt.teacher, tt.subjects, tt.submission, tt.lessons))
		})
	}
}

func TestTeacherSubmissionDetails(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name        string
		teachers    entity.Teachers
		subjects    map[string]entity.Subjects
		submissions map[string]*entity.TeacherSubmission
		lessons     map[string]entity.Lessons
		expect      TeacherSubmissionDetails
	}{
		{
			name: "success",
			teachers: entity.Teachers{
				{
					Teacher: &user.Teacher{
						Id:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "??????",
						FirstName:     "??????",
						LastNameKana:  "????????????",
						FirstNameKana: "????????????",
						Mail:          "teacher-test001@calmato.jp",
						Role:          user.Role_ROLE_TEACHER,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
				{
					Teacher: &user.Teacher{
						Id:            "teacherid",
						LastName:      "?????????",
						FirstName:     "??????",
						LastNameKana:  "?????????",
						FirstNameKana: "?????????",
						Mail:          "teacher-test002@calmato.jp",
						Role:          user.Role_ROLE_ADMINISTRATOR,
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
							Name:       "??????",
							Color:      "#F8BBD0",
							SchoolType: classroom.SchoolType_SCHOOL_TYPE_HIGH_SCHOOL,
							CreatedAt:  now.Unix(),
							UpdatedAt:  now.Unix(),
						},
					},
				},
			},
			submissions: map[string]*entity.TeacherSubmission{
				"teacherid": {
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			lessons: map[string]entity.Lessons{
				"teacherid": {
					{
						Lesson: &lesson.Lesson{
							Id:             1,
							ShiftSummaryId: 1,
							ShiftId:        1,
							SubjectId:      1,
							RoomId:         1,
							TeacherId:      "teacherid",
							StudentId:      "studentid",
							Notes:          "??????",
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
				},
			},
			expect: TeacherSubmissionDetails{
				{
					Teacher: &Teacher{
						ID:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "??????",
						FirstName:     "??????",
						LastNameKana:  "????????????",
						FirstNameKana: "????????????",
						Mail:          "teacher-test001@calmato.jp",
						Role:          RoleTeacher,
						CreatedAt:     now,
						UpdatedAt:     now,
						Subjects: map[SchoolType]Subjects{
							SchoolTypeElementarySchool: {},
							SchoolTypeJuniorHighSchool: {},
							SchoolTypeHighSchool: {
								{
									ID:         1,
									Name:       "??????",
									Color:      "#F8BBD0",
									SchoolType: SchoolTypeHighSchool,
									CreatedAt:  now,
									UpdatedAt:  now,
								},
							},
						},
					},
					IsSubmit:    false,
					LessonTotal: 0,
				},
				{
					Teacher: &Teacher{
						ID:            "teacherid",
						LastName:      "?????????",
						FirstName:     "??????",
						LastNameKana:  "?????????",
						FirstNameKana: "?????????",
						Mail:          "teacher-test002@calmato.jp",
						Role:          RoleAdministrator,
						CreatedAt:     now,
						UpdatedAt:     now,
						Subjects: map[SchoolType]Subjects{
							SchoolTypeElementarySchool: {},
							SchoolTypeJuniorHighSchool: {},
							SchoolTypeHighSchool:       {},
						},
					},
					IsSubmit:    true,
					LessonTotal: 1,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissionDetails(tt.teachers, tt.subjects, tt.submissions, tt.lessons))
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
