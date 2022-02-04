package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/internal/gateway/entity"
	"github.com/calmato/shs-web/api/proto/classroom"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
)

func TestStudentSchedules(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		schedules        entity.Schedules
		studentSchedules entity.StudentShiftSchedules
		expect           StudentSchedules
	}{
		{
			name: "success",
			schedules: entity.Schedules{
				{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Sunday),
						IsClosed: true,
						Lessons:  []*classroom.Schedule_Lesson{},
					},
				},
				{
					Schedule: &classroom.Schedule{
						Weekday:  int32(time.Monday),
						IsClosed: false,
						Lessons: []*classroom.Schedule_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
			},
			studentSchedules: entity.StudentShiftSchedules{
				{
					ShiftSchedule: &lesson.ShiftSchedule{
						Weekday: int32(time.Monday),
						Lessons: []*lesson.LessonSchedule{
							{StartTime: "1700", EndTime: "1830"},
						},
					},
				},
			},
			expect: StudentSchedules{
				{
					Weekday: time.Sunday,
					Lessons: StudentScheduleLessons{},
				},
				{
					Weekday: time.Monday,
					Lessons: StudentScheduleLessons{
						{StartTime: "1700", EndTime: "1830", Enabled: true},
						{StartTime: "1830", EndTime: "2000", Enabled: false},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			actual := NewStudentSchedules(tt.schedules, tt.studentSchedules)
			assert.Equal(t, tt.expect, actual)
		})
	}
}

func TestStudentSuggestedLesson(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		lesson *entity.StudentSuggestedLesson
		expect *StudentSuggestedLesson
	}{
		{
			name: "success",
			lesson: &entity.StudentSuggestedLesson{
				SuggestedLesson: &lesson.SuggestedLesson{
					SubjectId: 1,
					Total:     4,
				},
			},
			expect: &StudentSuggestedLesson{
				SubjectID: 1,
				Total:     4,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSuggestedLesson(tt.lesson))
		})
	}
}

func TestStudentSuggestedLessons(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		lessons entity.StudentSuggestedLessons
		expect  StudentSuggestedLessons
	}{
		{
			name: "success",
			lessons: entity.StudentSuggestedLessons{
				{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 1, Total: 4}},
				{SuggestedLesson: &lesson.SuggestedLesson{SubjectId: 2, Total: 4}},
			},
			expect: StudentSuggestedLessons{
				{SubjectID: 1, Total: 4},
				{SubjectID: 2, Total: 4},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentSuggestedLessons(tt.lessons))
		})
	}
}
