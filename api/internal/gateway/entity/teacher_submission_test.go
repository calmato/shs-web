package entity

import (
	"testing"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestTeacherSubmission(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name       string
		submission *lesson.TeacherSubmission
		expect     *TeacherSubmission
	}{
		{
			name: "success",
			submission: &lesson.TeacherSubmission{
				TeacherId:      "teacherid",
				ShiftSummaryId: 1,
				Decided:        true,
				CreatedAt:      now.Unix(),
				UpdatedAt:      now.Unix(),
			},
			expect: &TeacherSubmission{
				TeacherSubmission: &lesson.TeacherSubmission{
					TeacherId:      "teacherid",
					ShiftSummaryId: 1,
					Decided:        true,
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
			assert.Equal(t, tt.expect, NewTeacherSubmission(tt.submission))
		})
	}
}

func TestTeacherSubmissions(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions []*lesson.TeacherSubmission
		expect      TeacherSubmissions
	}{
		{
			name: "success",
			submissions: []*lesson.TeacherSubmission{
				{
					TeacherId:      "teacherid",
					ShiftSummaryId: 1,
					Decided:        true,
					CreatedAt:      now.Unix(),
					UpdatedAt:      now.Unix(),
				},
			},
			expect: TeacherSubmissions{
				{
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewTeacherSubmissions(tt.submissions))
		})
	}
}

func TestTeacherSubmissions_MapByShiftSummaryID(t *testing.T) {
	t.Parallel()
	now := jst.Now()
	tests := []struct {
		name        string
		submissions TeacherSubmissions
		expect      map[int64]*TeacherSubmission
	}{
		{
			name: "success",
			submissions: TeacherSubmissions{
				{
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				{
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 2,
						Decided:        false,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
			},
			expect: map[int64]*TeacherSubmission{
				1: {
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 1,
						Decided:        true,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
					},
				},
				2: {
					TeacherSubmission: &lesson.TeacherSubmission{
						TeacherId:      "teacherid",
						ShiftSummaryId: 2,
						Decided:        false,
						CreatedAt:      now.Unix(),
						UpdatedAt:      now.Unix(),
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
