package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentShiftTemplate(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		template *lesson.StudentShiftTemplate
		expect   *StudentShiftTemplate
	}{
		{
			name: "success",
			template: &lesson.StudentShiftTemplate{
				StudentId: "studentid",
				Schedules: []*lesson.ShiftSchedule{
					{
						Weekday: int32(time.Sunday),
						Lessons: []*lesson.LessonSchedule{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SuggestedLessons: []*lesson.SuggestedLesson{
					{SubjectId: 1, Total: 4},
				},
				CreatedAt: now.Unix(),
				UpdatedAt: now.Unix(),
			},
			expect: &StudentShiftTemplate{
				StudentID: "studentid",
				Schedules: StudentShiftSchedules{
					{
						ShiftSchedule: &lesson.ShiftSchedule{
							Weekday: int32(time.Sunday),
							Lessons: []*lesson.LessonSchedule{
								{StartTime: "1700", EndTime: "1830"},
								{StartTime: "1830", EndTime: "2000"},
							},
						},
					},
				},
				Lessons: StudentSuggestedLessons{
					{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
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
			actual := NewStudentShiftTemplate(tt.template)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudentShiftSchedules_LessonKeys(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		schedules StudentShiftSchedules
		expect    []string
	}{
		{
			name: "success",
			schedules: StudentShiftSchedules{
				{
					ShiftSchedule: &lesson.ShiftSchedule{
						Weekday: int32(time.Sunday),
						Lessons: []*lesson.LessonSchedule{
							{StartTime: "1530", EndTime: "1700"},
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				{
					ShiftSchedule: &lesson.ShiftSchedule{
						Weekday: int32(time.Monday),
						Lessons: []*lesson.LessonSchedule{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
			},
			expect: []string{
				"0:1530:1700",
				"0:1700:1830",
				"0:1830:2000",
				"1:1700:1830",
				"1:1830:2000",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := tt.schedules.LessonKeys()
			assert.ElementsMatch(t, tt.expect, actual.Strings())
		})
	}
}
