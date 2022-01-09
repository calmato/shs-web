package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/pkg/jst"
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
		name    string
		teacher *entity.Teacher
		shifts  entity.TeacherShifts
		expect  *TeacherSubmissionDetail
	}{
		{
			name: "success",
			teacher: &entity.Teacher{
				Teacher: &user.Teacher{
					Id:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Role:          user.Role_ROLE_TEACHER,
					CreatedAt:     now.Unix(),
					UpdatedAt:     now.Unix(),
				},
			},
			shifts: entity.TeacherShifts{
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "kSByoE6FetnPs5Byk3a9Zx",
						ShiftId:        1,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				{
					TeacherShift: &lesson.TeacherShift{
						TeacherId:      "kSByoE6FetnPs5Byk3a9Zx",
						ShiftId:        3,
						ShiftSummaryId: 1,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: &TeacherSubmissionDetail{
				Teacher: &Teacher{
					ID:            "kSByoE6FetnPs5Byk3a9Zx",
					LastName:      "中村",
					FirstName:     "広大",
					LastNameKana:  "なかむら",
					FirstNameKana: "こうだい",
					Mail:          "teacher-test001@calmato.jp",
					Role:          RoleTeacher,
					CreatedAt:     now,
					UpdatedAt:     now,
				},
				LessonTotal: 2,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissionDetail(tt.teacher, tt.shifts))
		})
	}
}

func TestTeacherSubmissionDetails(t *testing.T) {
	t.Parallel()
	now := jst.Date(2021, 8, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		teachers entity.Teachers
		shifts   map[string]entity.TeacherShifts
		expect   TeacherSubmissionDetails
	}{
		{
			name: "success",
			teachers: entity.Teachers{
				{
					Teacher: &user.Teacher{
						Id:            "kSByoE6FetnPs5Byk3a9Zx",
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          user.Role_ROLE_TEACHER,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
				{
					Teacher: &user.Teacher{
						Id:            "teacherid",
						LastName:      "テスト",
						FirstName:     "講師",
						LastNameKana:  "てすと",
						FirstNameKana: "こうし",
						Mail:          "teacher-test002@calmato.jp",
						Role:          user.Role_ROLE_ADMINISTRATOR,
						CreatedAt:     now.Unix(),
						UpdatedAt:     now.Unix(),
					},
				},
			},
			shifts: map[string]entity.TeacherShifts{
				"kSByoE6FetnPs5Byk3a9Zx": {
					{
						TeacherShift: &lesson.TeacherShift{
							TeacherId:      "kSByoE6FetnPs5Byk3a9Zx",
							ShiftId:        1,
							ShiftSummaryId: 1,
							CreatedAt:      now.Unix(),
							UpdatedAt:      now.Unix(),
						},
					},
					{
						TeacherShift: &lesson.TeacherShift{
							TeacherId:      "kSByoE6FetnPs5Byk3a9Zx",
							ShiftId:        3,
							ShiftSummaryId: 1,
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
						LastName:      "中村",
						FirstName:     "広大",
						LastNameKana:  "なかむら",
						FirstNameKana: "こうだい",
						Mail:          "teacher-test001@calmato.jp",
						Role:          RoleTeacher,
						CreatedAt:     now,
						UpdatedAt:     now,
					},
					LessonTotal: 2,
				},
				{
					Teacher: &Teacher{
						ID:            "teacherid",
						LastName:      "テスト",
						FirstName:     "講師",
						LastNameKana:  "てすと",
						FirstNameKana: "こうし",
						Mail:          "teacher-test002@calmato.jp",
						Role:          RoleAdministrator,
						CreatedAt:     now,
						UpdatedAt:     now,
					},
					LessonTotal: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissionDetails(tt.teachers, tt.shifts))
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
