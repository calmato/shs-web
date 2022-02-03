package entity

import (
	"testing"
	"time"

	"github.com/calmato/shs-web/api/pkg/jst"
	"github.com/calmato/shs-web/api/proto/lesson"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestStudentShiftTemplate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		studentID string
		template  *lesson.StudentShiftTemplateToUpsert
		expect    *StudentShiftTemplate
	}{
		{
			name:      "success",
			studentID: "studentid",
			template: &lesson.StudentShiftTemplateToUpsert{
				Schedules: []*lesson.StudentShiftTemplateToUpsert_Schedule{
					{
						Weekday: int32(time.Sunday),
						Lessons: []*lesson.StudentShiftTemplateToUpsert_Lesson{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SuggestedLessons: []*lesson.StudentSuggestedLesson{
					{SubjectId: 1, Total: 4},
				},
			},
			expect: &StudentShiftTemplate{
				StudentID: "studentid",
				Schedules: ShiftSchedules{
					{
						Weekday: time.Sunday,
						Lessons: LessonSchedules{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expect, NewStudentShiftTemplate(tt.studentID, tt.template))
		})
	}
}

func TestStudentShiftTemplate_Fill(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		template *StudentShiftTemplate
		expect   *StudentShiftTemplate
		isErr    bool
	}{
		{
			name: "success",
			template: &StudentShiftTemplate{
				StudentID:            "studentid",
				SchedulesJSON:        datatypes.JSON([]byte(`[{"weekday":0,"lessons":[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"2000"}]}]`)),
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4}]`)),
			},
			expect: &StudentShiftTemplate{
				StudentID: "studentid",
				Schedules: ShiftSchedules{
					{
						Weekday: time.Sunday,
						Lessons: LessonSchedules{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SchedulesJSON: datatypes.JSON([]byte(`[{"weekday":0,"lessons":[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"2000"}]}]`)),
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4}]`)),
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.template.Fill()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.template)
		})
	}
}

func TestStudentShiftTemplate_FillJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		template *StudentShiftTemplate
		expect   *StudentShiftTemplate
		isErr    bool
	}{
		{
			name: "success",
			template: &StudentShiftTemplate{
				StudentID: "studentid",
				Schedules: ShiftSchedules{
					{
						Weekday: time.Sunday,
						Lessons: LessonSchedules{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
			},
			expect: &StudentShiftTemplate{
				StudentID: "studentid",
				Schedules: ShiftSchedules{
					{
						Weekday: time.Sunday,
						Lessons: LessonSchedules{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SchedulesJSON: datatypes.JSON([]byte(`[{"weekday":0,"lessons":[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"2000"}]}]`)),
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4}]`)),
			},
			isErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.template.FillJSON()
			assert.Equal(t, tt.isErr, err != nil, err)
			assert.Equal(t, tt.expect, tt.template)
		})
	}
}

func TestStudentShiftTemplate_Proto(t *testing.T) {
	t.Parallel()
	now := jst.Date(2022, 1, 2, 18, 30, 0, 0)
	tests := []struct {
		name     string
		template *StudentShiftTemplate
		expect   *lesson.StudentShiftTemplate
	}{
		{
			name: "success",
			template: &StudentShiftTemplate{
				StudentID: "studentid",
				Schedules: ShiftSchedules{
					{
						Weekday: time.Sunday,
						Lessons: LessonSchedules{
							{StartTime: "1700", EndTime: "1830"},
							{StartTime: "1830", EndTime: "2000"},
						},
					},
				},
				SchedulesJSON: datatypes.JSON([]byte(`[{"weekday":0,"lessons":[{"startTime":"1700","endTime":"1830"},{"startTime":"1830","endTime":"2000"}]}]`)),
				SuggestedLessons: SuggestedLessons{
					{SubjectID: 1, Total: 4},
				},
				SuggestedLessonsJSON: datatypes.JSON([]byte(`[{"subjectId":1,"total":4}]`)),
				CreatedAt:            now,
				UpdatedAt:            now,
			},
			expect: &lesson.StudentShiftTemplate{
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
				SuggesteLessons: []*lesson.SuggestedLesson{
					{SubjectId: 1, Total: 4},
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
			assert.Equal(t, tt.expect, tt.template.Proto())
		})
	}
}
